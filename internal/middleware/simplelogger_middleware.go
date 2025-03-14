package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SimpleLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {

		fmt.Printf("[%s] %s requested\n", c.Method(), c.OriginalURL())
		return c.Next()
	}
}