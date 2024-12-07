package product

import (
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
		productRouter.Post("", handler.CreateProduct)
		productRouter.Get("/sku/:sku", handler.GetProductDetail)
	}
}