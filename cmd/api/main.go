package main

import (
	"log"

	"github.com/ardwiinoo/online-shop/external/database"
	"github.com/ardwiinoo/online-shop/internal/config"
)

func main() {
	fileName := "cmd/api/config.yaml"

	if err := config.LoadConfig(fileName); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}
} 