package infrafiber

import (
	"log"
	"strings"

	"github.com/ardwiinoo/online-shop/infra/response"
	"github.com/ardwiinoo/online-shop/internal/config"
	"github.com/ardwiinoo/online-shop/utility"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")

		if authorization == "" {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(ctx)
		}

		bearer := strings.Split(authorization, "Bearer ")

		if len(bearer) != 2 {
			log.Println("token invalid")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(ctx)
		}

		token := bearer[1]

		publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)

		if err != nil {
			log.Println("token invalid")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(ctx)
		}

		ctx.Locals("PUBLIC_ID", publicId)
		ctx.Locals("ROLE", role)

		return ctx.Next()
	}
}