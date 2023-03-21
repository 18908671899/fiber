package controller

import (
	"fiber/models"
	"fiber/pkg/logging"
	"fiber/pkg/validates"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	login := &models.Login{}
	code := 200
	message := "SUCCESS"

	if err := c.BodyParser(login); err != nil {
		code = 500
		message = "参数解析错误"
		logging.Error(err)
	}

	errors := validates.ValidateStruct(*login)

	if errors != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    code,
			"message": message,
			"data":    errors,
		})
	}

	res := models.GetToken(login)

	if res == "" {
		code = fiber.StatusUnauthorized
		message = "账户或密码错误"
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    code,
			"message": message,
			"data":    res,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    code,
		"message": message,
		"data":    res,
	})
}
