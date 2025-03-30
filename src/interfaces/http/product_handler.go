package interfaces

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/johancuervo/apiServiceGo/src/application"
	"github.com/johancuervo/apiServiceGo/src/domain"
)

// ProductHandler maneja los endpoints relacionados con productos
type ProductHandler struct {
	productUseCase *application.ProductUseCase
}

// NewProductHandler crea un nuevo manejador de productos
func NewProductHandler(useCase *application.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: useCase}
}

// HandleSaveProducts procesa la solicitud para guardar productos
func (h *ProductHandler) HandleSaveProducts(c *fiber.Ctx) error {
	var products []domain.Product

	if err := c.BodyParser(&products); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	if len(products) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No products to save"})
	}

	if err := h.productUseCase.SaveProducts(products); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save products"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Products stored successfully"})
}

// NewProxyHandler crea un nuevo ProxyHandler
// HandleGetProducts godoc
// @Summary Obtiene una lista de productos
// @Description Retorna productos filtrados por categoría y paginados
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /products [get]
func (h *ProductHandler) HandleGetProducts(c *fiber.Ctx) error {
	products, err := h.productUseCase.GetProducts()
	if err != nil {
		log.Printf("❌ Error al obtener productos: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch products",
		})
	}

	return c.JSON(products)
}
