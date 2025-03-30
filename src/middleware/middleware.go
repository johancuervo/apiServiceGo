package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})
}
func Cors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://example.com, https://otro-dominio.com", // Dominios permitidos
		AllowMethods:     "GET,POST,PUT,DELETE",                           // Métodos permitidos
		AllowHeaders:     "Content-Type, Authorization",                   // Headers permitidos
		AllowCredentials: true,                                            // Permitir credenciales (cookies, auth headers)
	}))
}
