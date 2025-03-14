package routes

import (
	"fiber-test/internal/handlers"
	"fiber-test/internal/repository"
	"fiber-test/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userApi := app.Group("/users")

	userApi.Get("/", userHandler.GetAllUsers)
	userApi.Get("/:id", userHandler.GetUserById)
	
	userApi.Post("/", userHandler.CreateUser)
	userApi.Put("/:id", userHandler.UpdateUser)
	userApi.Delete("/:id", userHandler.DeleteUser)
}