package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productsRouter := router.Group("products")
	{
		productsRouter.Get("", handler.GetListProducts)
		productsRouter.Post("", handler.CreateProduct)
		productsRouter.Get("/sku/:sku", handler.GetProductDetail)
	}
}