package main

import (
	"log"

	"github.com/ardwiinoo/online-shop/apps/auth"
	"github.com/ardwiinoo/online-shop/apps/product"
	"github.com/ardwiinoo/online-shop/apps/transaction"
	"github.com/ardwiinoo/online-shop/external/database"
	"github.com/ardwiinoo/online-shop/internal/config"
	"github.com/gofiber/fiber/v2"
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

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)
	product.Init(router, db)
	transaction.Init(router, db)

	router.Listen(config.Cfg.App.Port)
} 