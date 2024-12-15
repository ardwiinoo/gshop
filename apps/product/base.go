package product

import (
	"github.com/ardwiinoo/online-shop/apps/auth"
	infrafiber "github.com/ardwiinoo/online-shop/infra/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRouter := router.Group("products")
	{
		productRouter.Get("", handler.GetListProducts)
		productRouter.Get("/sku/:sku", handler.GetProductDetail)

		productRouter.Post("", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_ADMIN)}), handler.CreateProduct)
	}
}