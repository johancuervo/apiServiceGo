package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/johancuervo/apiServiceGo/src/application"
	"github.com/johancuervo/apiServiceGo/src/domain"
	infrastructure "github.com/johancuervo/apiServiceGo/src/infraestructure/external"
)

// ProxyHandler maneja la lógica del proxy con Vercel
type ProxyHandler struct {
	authManager    *infrastructure.AuthManager
	productUseCase *application.ProductUseCase
}

// HandleProxyRequest godoc
// @Summary Redirecciona una solicitud a Vercel
// @Description Proxy que reenvía solicitudes a la API de Vercel
// @Tags Proxy
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /proxy/vercel/api/Productos/Stock [get]
func NewProxyHandler(auth *infrastructure.AuthManager, productUC *application.ProductUseCase) *ProxyHandler {
	return &ProxyHandler{authManager: auth, productUseCase: productUC}
}

// HandleProxyRequest maneja las solicitudes a Vercel y guarda productos si es necesario
func (h *ProxyHandler) HandleProxyRequest(c *fiber.Ctx) error {
	// Verificamos si el token es válido
	if !h.authManager.IsTokenValid() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No valid authentication token available",
		})
	}

	// Construimos la URL destino
	path := c.Params("*")
	baseURL := "https://sil-recibo-despacho-nodejs-git-desarrollo-johancuervos-projects.vercel.app"
	targetURL := baseURL + "/" + path

	// Hacemos la solicitud a Vercel
	client := &http.Client{}
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating request"})
	}

	// Añadimos el token de Vercel en la cabecera
	req.Header.Add("Cookie", "_vercel_jwt="+h.authManager.GetToken())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Referer", baseURL+"/swagger/")
	req.Header.Add("User-Agent", "GoFiber/2.0")

	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error making request to Vercel",
			"details": err.Error(),
		})
	}
	defer resp.Body.Close()

	// Leemos la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error reading response"})
	}

	// Si es un endpoint de stock y la respuesta es 200, guardamos en MongoDB
	if resp.StatusCode == 200 && isStockEndpoint(path) {
		var products []domain.Product
		if err := json.Unmarshal(body, &products); err != nil {
			// Intentamos parsearlo como un solo producto
			var singleProduct domain.Product
			if errSingle := json.Unmarshal(body, &singleProduct); errSingle == nil {
				products = []domain.Product{singleProduct}
			} else {
				log.Printf("Warning: Could not parse response as product(s): %v", err)
			}
		}

		// Guardamos productos si existen
		if len(products) > 0 {
			err := h.productUseCase.SaveProducts(products)
			if err != nil {
				log.Printf("Error saving products to MongoDB: %v", err)
			}
		}
	}

	// Respondemos con el mismo código de estado
	c.Status(resp.StatusCode)

	// Copiamos los headers relevantes
	for key, values := range resp.Header {
		for _, value := range values {
			c.Append(key, value)
		}
	}

	return c.Send(body)
}

// isStockEndpoint verifica si la ruta pertenece a stock
func isStockEndpoint(path string) bool {
	fmt.Println(path)
	return path == "api/Productos/Stock" || path == "inventory"
}
