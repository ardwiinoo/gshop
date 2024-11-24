package products

import (
	"testing"

	"github.com/ardwiinoo/online-shop/infra/response"
	"github.com/stretchr/testify/require"
)

func TestValidateProduct(t *testing.T) {
	t.Run("Success validate product", func(t *testing.T) {
		product := Product{
			Name: "Product 1",
			Stock: 10,
			Price: 10_000,
		}

		err := product.Validate()
		require.Nil(t, err)
	})

	t.Run("Failed validate product, name is required", func(t *testing.T) {
		product := Product{
			Name: "",
			Stock: 10,
			Price: 10_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})

	t.Run("Failed validate product, name is invalid", func(t *testing.T) {
		product := Product{
			Name: "P1",
			Stock: 10,
			Price: 10_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductNameInvalid, err)
	})

	t.Run("Failed validate product, stock is invalid", func(t *testing.T) {
		product := Product{
			Name: "Product 1",
			Stock: 0,
			Price: 10_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)
	})

	t.Run("Failed validate product, price is invalid", func(t *testing.T) {
		product := Product{
			Name: "Product 1",
			Stock: 10,
			Price: 0,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}