package main

import (
	"fiber-test/internal/database"
	"fiber-test/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()

	app := fiber.New(fiber.Config{
		AppName: "Fiber test API",
		ErrorHandler: customErrorHandler,
	})

	routes.SetupRoutes(app)
	
	port := os.Getenv("PORT")
	if port == "" { port = "3000" }
	log.Fatal(app.Listen(":" + port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}