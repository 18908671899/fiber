package routers

import (
	"fiber/controller"
	"fiber/middleware/jwt"
	"fiber/pkg/logging"
	"fiber/pkg/setting"
	"fiber/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	jwtware "github.com/gofiber/jwt/v3"
)

func InitRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
	})

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[INFO-${locals:requestid}]${time} pid: ${pid} status:${status} - ${method} path: ${path} queryParams: ${queryParams} body: ${body}\n resBody: ${resBody}\n error: ${error}\n",
		Output: logging.F,
	}))
	apiv1 := app.Group("/v1")

	{
		app.Post("/login", controller.Login)
	}

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: utils.JwtError,
	}))

	apiv1.Use(jwt.Jwt)

	{
		apiv1.Get("/test", controller.GetTest)
		apiv1.Post("/test", controller.AddTest)
		apiv1.Put("/test/:id", controller.EditTest)
		apiv1.Delete("/test/:id", controller.DelTest)
	}
	return app
}
