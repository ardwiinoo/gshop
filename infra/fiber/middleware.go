package infrafiber

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ardwiinoo/online-shop/infra/response"
	"github.com/ardwiinoo/online-shop/internal/config"
	"github.com/ardwiinoo/online-shop/utility"
	"github.com/gofiber/fiber/v2"
)

func Trace() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		err := ctx.Next()

		end := time.Now()
		latency := end.Sub(start)

		method := ctx.Method()
		path := ctx.Path()
		status := ctx.Response().StatusCode()
		ip := ctx.IP()

		log.Printf("%s - %s %s %d %s", ip, method, path, status, latency)

		return err
	}
}

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

func CheckRoles(authorizedRoles []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		role := fmt.Sprintf("%v", ctx.Locals("ROLE"))

		isExists := false
		for _, authauthorizedRole := range authorizedRoles {
			if role == authauthorizedRole {
				isExists = true
				break
			}
		}

		if !isExists {
			return NewResponse(
				WithError(response.ErrorForbiddenAccess),
			).Send(ctx)
		}

		return ctx.Next()
	}
}