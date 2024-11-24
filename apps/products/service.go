package products

import (
	"context"

	"github.com/ardwiinoo/online-shop/infra/response"
)

type Repository interface{
	CreateProduct(ctx context.Context, model Product) (err error)
	GetAllProductsWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateProduct(ctx context.Context, req CreateProductRequestPayload) (err error) {
	productEntity := NewProductFromCreateProductRequest(req)

	if err = productEntity.Validate(); err != nil {
		return
	}

	if err = s.repo.CreateProduct(ctx, productEntity); err != nil {
		return
	}

	return
}

func (s service) ListProducts(ctx context.Context, req ListProductRequestPayload) (products []Product, err error) {
	pagination := NewProductPaginationFromListProductRequest(req)

	products, err = s.repo.GetAllProductsWithPaginationCursor(ctx, pagination)

	if err != nil {
		if err == response.ErrNotFound {
			return []Product{}, nil
		}

		return
	}

	if len(products) == 0 {
		return []Product{}, nil
	}

	return
}