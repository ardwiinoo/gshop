package transaction

import (
	"context"
	"testing"

	"github.com/ardwiinoo/online-shop/external/database"
	"github.com/ardwiinoo/online-shop/internal/config"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)

	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)

	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateTransaction(t *testing.T) {
	t.Run("Success Create Transaction", func(t *testing.T) {
		req := CreateTransactionRequestPayload {
			ProductSKU: "cad5c7e1-e7f7-4ee6-b8e5-65191a5e596b",
			Amount: 2,
			UserPublicId: "9cf6864c-5a89-4847-a6d2-be3cb4d9dd55",
		}

		err := svc.CreateTransaction(context.Background(), req)
		require.Nil(t, err)
	})
}