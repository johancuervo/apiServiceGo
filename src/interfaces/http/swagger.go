package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// SetupSwagger configura Swagger en Fiber
func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // Ruta de Swagger
}
