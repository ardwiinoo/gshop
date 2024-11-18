package auth

import (
	"testing"

	"github.com/ardwiinoo/online-shop/infra/response"
	"github.com/stretchr/testify/require"
)

func TestValidateAuthEntity(t *testing.T) {
	t.Run("Success Validate AuthEntity", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "ardwiinoo@gmail.com",
			Password: "password",
		}

		err := authEntity.Validate()
		require.Nil(t, err)
	})
	
	t.Run("Failed Validate, email is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "",
			Password: "password",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})

	t.Run("Failed Validate, email is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "ardwiinoogmail.com",
			Password: "password",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})

	t.Run("Failed Validate, password is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "ardwiinoo@gmail.com",
			Password: "",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})

	t.Run("Failed Validate, password must have min 8 chars", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "ardwiinoo@gmail.com",
			Password: "pass",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordLength, err)
	})
}