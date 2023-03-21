package jwt

import (
	"fiber/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Jwt(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	logging.Info(claims)
	return ctx.Next()
}
