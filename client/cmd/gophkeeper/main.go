package main

import (
	"log"

	"github.com/iurnickita/gophkeeper/client/internal/cache"
	"github.com/iurnickita/gophkeeper/client/internal/cli"
	"github.com/iurnickita/gophkeeper/client/internal/config"
	grpcclient "github.com/iurnickita/gophkeeper/client/internal/grpc_client/client"
	"github.com/iurnickita/gophkeeper/client/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.GetConfig()

	// Клиент
	client, err := grpcclient.NewClient(cfg.GRPCClient)
	if err != nil {
		return err
	}

	// Кэш
	cache, err := cache.NewCache(cfg.Cache)
	if err != nil {
		return err
	}

	// Логика
	service, err := service.NewService(cfg.Service, client, cache)
	if err != nil {
		return err
	}

	// Пользовательский интерфейс
	cli.Execute(service)

	// Завершение работы
	service.Close()
	return nil
}
