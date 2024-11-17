package database

import (
	"testing"

	"github.com/ardwiinoo/online-shop/internal/config"
	"github.com/stretchr/testify/require"
)

func init() {
	fileName := "../../cmd/api/config.yaml"

	err := config.LoadConfig(fileName)
	if err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("Success Connect Postgres", func(t *testing.T) {
		db, err := ConnectPostgres(config.Cfg.DB)

		require.Nil(t, err)
		require.NotNil(t, db)
	})

	t.Run("Failed Connect Postgres", func(t *testing.T) {
		cfg := config.Cfg.DB
		cfg.Password = "invalid password"

		db, err := ConnectPostgres(cfg)

		require.NotNil(t, err)
		require.Nil(t, db)
	})
}