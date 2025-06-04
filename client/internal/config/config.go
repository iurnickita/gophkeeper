// Пакет config. Конфигурация с помощью флагов/переменных среды/значений по умолчанию
package config

import (
	cacheConfig "github.com/iurnickita/gophkeeper/client/internal/cache/config"
	grpcClientConig "github.com/iurnickita/gophkeeper/client/internal/grpc_client/client/config"
	serviceConfig "github.com/iurnickita/gophkeeper/client/internal/service/config"
)

// Config - общая конфигурация
type Config struct {
	GRPCClient grpcClientConig.Config
	Service    serviceConfig.Config
	Cache      cacheConfig.Config
}

// GetConfig собирает конфигурацию сервиса
func GetConfig() Config {
	cfg := Config{}

	if cfg.Cache.ValidPeriod == 0 {
		cfg.Cache.ValidPeriod = 1
	}

	return cfg
}
