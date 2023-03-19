package routers

import (
	"fiber/controller"
	"fiber/pkg/setting"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
	})

	app.Use(requestid.New())

	apiv1 := app.Group("/v1")

	{
		apiv1.Get("/test", controller.GetTest)
		apiv1.Post("/test", controller.AddTest)
		apiv1.Put("/test/:id", controller.EditTest)
		apiv1.Delete("/test/:id", controller.DelTest)
	}
	return app
}
