package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johancuervo/apiServiceGo/src/application"
	"github.com/johancuervo/apiServiceGo/src/models"
)

// AuthHandler maneja las solicitudes HTTP de autenticación
type AuthHandler struct {
	authUseCase *application.AuthUseCase
}

// NewAuthHandler crea un nuevo manejador
func NewAuthHandler(authUseCase *application.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

// HandleSetVercelToken godoc
// @Summary Autentica con Vercel
// @Description Obtiene un token de autenticación de Vercel con email y contraseña
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.AuthRequest true "Datos de autenticación"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/vercel [post]
func (h *AuthHandler) HandleSetVercelToken(c *fiber.Ctx) error {
	// Estructura para recibir el token

	authReq := new(models.AuthRequest)
	if err := c.BodyParser(authReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if authReq.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token is required",
		})
	}

	// Guardamos el token usando la capa de aplicación
	h.authUseCase.RefreshToken(authReq.Token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Vercel authentication token stored successfully",
	})
}
