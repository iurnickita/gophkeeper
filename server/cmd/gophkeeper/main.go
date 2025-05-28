package main

import (
	"log"

	"github.com/iurnickita/gophkeeper/server/internal/auth"
	"github.com/iurnickita/gophkeeper/server/internal/config"
	grpcserver "github.com/iurnickita/gophkeeper/server/internal/grpc_server/server"
	"github.com/iurnickita/gophkeeper/server/internal/logger"
	"github.com/iurnickita/gophkeeper/server/internal/service"
	"github.com/iurnickita/gophkeeper/server/internal/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.GetConfig()

	zaplog, err := logger.NewZapLog(cfg.Logger)
	if err != nil {
		return err
	}

	store, err := store.NewStore(cfg.Store)
	if err != nil {
		return err
	}

	auth, err := auth.NewAuth(store)
	if err != nil {
		return err
	}

	service, err := service.NewService(cfg.Service, store, zaplog)
	if err != nil {
		return err
	}

	return grpcserver.Serve(cfg.GRPCServer, auth, service, zaplog)
}
