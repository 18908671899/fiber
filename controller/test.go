package controller

import "github.com/gofiber/fiber/v2"

func GetTest(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "test",
	})
}
