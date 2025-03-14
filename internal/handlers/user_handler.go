package handlers

import (
	"fiber-test/internal/models"
	"fiber-test/internal/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler {
		service: service,
	}
}


func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers() 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error during get users",
		})
	}
	return c.JSON(users)
}


func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid form data for get user",
		})
	}
	
	user, err := h.service.GetUserById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf(`User %d does not exist`, id),
		})
	}
	return c.JSON(user)
}


func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid form data for post user",
		})
	}

	createdUser, err := h.service.CreateUser(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unknown error during post user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf(`User %d crée avec succès`, createdUser.ID),
		"user": createdUser,
	})
}


func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid form data for get user",
		})
	}

	var updatedUser models.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid form data for put user", 
		})
	}

	user, err := h.service.UpdateUser(uint(id), updatedUser)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf(`User %d not found on put user`, id),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf(`User %d successfully updated`, id),
		"user": user,
	})
}


func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid form data for delete user",
		})
	}

	err = h.service.DeleteUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf(`User %d not found on delete user`, id),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf(`User %d successfully deleted`, id),
	})
}