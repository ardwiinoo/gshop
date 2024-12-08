package transaction

import (
	"context"

	"github.com/ardwiinoo/online-shop/infra/response"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	TransactionRepository
	ProductRepository
	TransactionDBRepository
}

type TransactionRepository interface {
	CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx Transaction) (err error)
}
type ProductRepository interface{
	GetProductBySKU(ctx context.Context, productSKU string) (model Product, err error)
	UpdateProductStockWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error)
}

type TransactionDBRepository interface {
	Begin(ctx context.Context) (tx *sqlx.Tx, err error)
	Commit(ctx context.Context, tx *sqlx.Tx) (err error)
	Rollback(ctx context.Context, tx *sqlx.Tx) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateTransaction(ctx context.Context, req CreateTransactionRequestPayload) (err error) {
	myProduct, err := s.repo.GetProductBySKU(ctx, req.ProductSKU)

	if err != nil {
		return
	}

	if !myProduct.IsExists() {
		err = response.ErrNotFound
		return
	}

	trx := NewTransactionFromCreateRequest(req)
	trx.FromProduct(myProduct).SetPlatformFee(1_000).SetGrandTotal()

	if err = trx.Validate(); err != nil {
		return
	}

	if err = trx.ValidateStock(uint8(myProduct.Stock)); err != nil {
		return
	}

	tx, err := s.repo.Begin(ctx)

	if err != nil {
		return
	}

	defer s.repo.Rollback(ctx, tx)

	if err = s.repo.CreateTransactionWithTx(ctx, tx,trx); err != nil {
		return
	}

	if err = myProduct.UpdateStockProduct(trx.Amount); err != nil {
		return
	}

	if err = s.repo.UpdateProductStockWithTx(ctx, tx, myProduct); err != nil {
		return
	}

	if err = s.repo.Commit(ctx, tx); err != nil {
		return
	}

	return
}
