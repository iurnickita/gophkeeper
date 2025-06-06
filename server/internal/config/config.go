// Пакет config. Конфигурация с помощью флагов/переменных среды/значений по умолчанию
package config

import (
	"flag"
	"os"

	grpcServerConig "github.com/iurnickita/gophkeeper/server/internal/grpc_server/server/config"
	loggerConfig "github.com/iurnickita/gophkeeper/server/internal/logger/config"
	serviceConfig "github.com/iurnickita/gophkeeper/server/internal/service/config"
	storeConfig "github.com/iurnickita/gophkeeper/server/internal/store/config"
)

// Config - общая конфигурация
type Config struct {
	GRPCServer grpcServerConig.Config
	Service    serviceConfig.Config
	Store      storeConfig.Config
	Logger     loggerConfig.Config
}

// GetConfig собирает конфигурацию сервиса
func GetConfig() Config {
	cfg := Config{}

	// Флаги
	flag.StringVar(&cfg.Store.DBDsn, "d", "", "database dsn")

	// Переменные окружения
	if envdsn := os.Getenv("DATABASE_URI"); envdsn != "" {
		cfg.Store.DBDsn = envdsn
	}

	// По умолчанию
	if cfg.Store.DBDsn == "" {
		cfg.Store.DBDsn = "host=localhost user=bob password=bob dbname=gophkeeper sslmode=disable"
	}

	return cfg
}
