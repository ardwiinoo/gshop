package auth

import (
	"context"
	"fmt"
	"testing"

	"github.com/ardwiinoo/online-shop/external/database"
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