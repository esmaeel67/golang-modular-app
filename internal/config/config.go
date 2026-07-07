package config

import (
	"os"
	"time"

	"github.com/esmaeel67/golang-modular-app/internal/rpc"
	"github.com/esmaeel67/golang-modular-app/internal/web"
	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

type PGConfig struct {
	// required:"true"
	Conn string `default:"postgres://postgres:admin@localhost:5432/mallbots?sslmode=disable&pool_max_conns=100&pool_min_conns=15"`
}

type AppConfig struct {
	Environment     string
	LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	PG              PGConfig
	Nats            NatsConfig
	Rpc             rpc.RpcConfig
	Web             web.WebConfig
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

type NatsConfig struct {
	URL    string `required:"true"`
	Stream string `default:"mallbots"`
}

func GetConfig() *AppConfig {
	return &AppConfig{}
}

func InitConfig() (cfg AppConfig, err error) {

	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)
	return
}
