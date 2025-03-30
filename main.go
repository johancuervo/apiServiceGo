package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/johancuervo/apiServiceGo/docs"
	"github.com/johancuervo/apiServiceGo/src/application"
	infrastructuredb "github.com/johancuervo/apiServiceGo/src/infraestructure/database"
	infraestructure "github.com/johancuervo/apiServiceGo/src/infraestructure/external"
	interfaces "github.com/johancuervo/apiServiceGo/src/interfaces/http"
	"github.com/johancuervo/apiServiceGo/src/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Configuración de MongoDB
	mongoRepo, err := infrastructuredb.NewMongoRepository()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Inicializar capas
	authManager := infraestructure.NewAuthManager()
	authUseCase := application.NewAuthUseCase(authManager)
	authHandler := interfaces.NewAuthHandler(authUseCase)
	productUseCase := application.NewProductUseCase(mongoRepo)
	productHandler := interfaces.NewProductHandler(productUseCase)
	proxyHandler := interfaces.NewProxyHandler(authManager, productUseCase)
	// Crear aplicación Fiber
	app := fiber.New()
	middleware.Cors(app)
	middleware.SetupRoutes(app)
	// Rutas de Swagger
	interfaces.SetupSwagger(app)
	// Definir rutas
	app.Post("/auth/vercel", authHandler.HandleSetVercelToken)
	app.Get("/proxy/vercel/*", proxyHandler.HandleProxyRequest)
	app.Get("/products", productHandler.HandleGetProducts)

	// Iniciar servidor
	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
