package config

import (
	"time"

	"github.com/esmaeel67/golang-modular-app/internal/rpc"
	"github.com/esmaeel67/golang-modular-app/internal/web"
)

type PGConfig struct {
	Conn string `required:"true"`
}

type AppConfig struct {
	Environment     string
	LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	PG              PGConfig
	Rpc             rpc.RpcConfig
	Web             web.WebConfig
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

func GetConfig() *AppConfig {
	return &AppConfig{}
}
