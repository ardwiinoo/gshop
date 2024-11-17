package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Success Load Config", func(t *testing.T) {
		fileName := "../../cmd/api/config.yaml"
		err := LoadConfig(fileName)

		require.Nil(t, err)
		log.Printf("%+v\n", Cfg)
	})

	t.Run("Failed Load Config", func(t *testing.T) {
		fileName := "config.yaml"
		err := LoadConfig(fileName)

		require.NotNil(t, err)
		log.Printf("%+v\n", Cfg)
	})
}