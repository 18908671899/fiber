package routers

import (
	"fiber/controller"
	"github.com/gofiber/fiber/v2"
)

func InitRouter() *fiber.App {
	app := fiber.New()

	apiv1 := app.Group("/v1")

	apiv1.Get("/test", controller.GetTest)

	return app
}
