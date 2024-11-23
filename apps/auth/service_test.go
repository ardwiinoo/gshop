package auth

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ardwiinoo/online-shop/external/database"
	"github.com/ardwiinoo/online-shop/infra/response"
	"github.com/ardwiinoo/online-shop/internal/config"
	"github.com/google/uuid"
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

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@gmail.id", uuid.NewString()),
		Password: "mysupersecretpassword",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}

func TestRegister_Fail(t *testing.T) {
	t.Run("Error email already used", func(t *testing.T) {
		// Prep
		email := fmt.Sprintf("%v@gmail.id", uuid.NewString())

		req := RegisterRequestPayload{
			Email:    email,
			Password: "mysupersecretpassword",
		}

		err := svc.register(context.Background(), req)
		require.Nil(t, err)
		// End Prep

		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlreadyUsed, err)
	})
}

func TestLogin_Succes(t *testing.T) {
	email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
	pass := "mysupersecretpassword"

	req := RegisterRequestPayload{
		Email: email,
		Password: pass,
	}

	err := svc.register(context.Background(), req)
	require.Nil(t, err)

	reqLogin := LoginRequestPayload{
		Email: email,
		Password: pass,
	}

	token, err := svc.login(context.Background(), reqLogin)
	require.Nil(t, err)
	require.NotEmpty(t, token)

	log.Println(token)

}