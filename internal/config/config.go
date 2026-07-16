package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/esmaeel67/golang-modular-app/internal/rpc"
	"github.com/esmaeel67/golang-modular-app/internal/web"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/stackus/dotenv"
)

type PGConfig struct {
	// required:"true"
	Conn string `default:"postgres://postgres:admin@localhost:5432/mallbots?sslmode=disable&pool_max_conns=100&pool_min_conns=15"`
}
type ServerConfig struct {
	InternalPort    string
	ExternalPort    string
	RunMode         string
	ShutdownTimeout time.Duration
}

type AppConfig struct {
	Server          ServerConfig
	Environment     string
	LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	Postgres        PostgresConfig
	Goose           GooseConfig
	Nats            NatsConfig
	Rpc             rpc.RpcConfig
	Web             web.WebConfig
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

type NatsConfig struct {
	Host   string `validate:"required"`
	Port   string `mapstructure:"port" validate:"required"`
	Stream string `validate:"required"`
}

func (nat NatsConfig) GetUrl() string {
	return fmt.Sprintf("%s:%s", nat.Host, nat.Port)
}

func GetConfig() *AppConfig {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error in load config %v", err)
	}
	cfg, err := ParseConfig(v)
	envPort := os.Getenv("PORT")
	log.Printf("envPort -> %s, cfgPath: %s", envPort, cfgPath)
	if envPort != "" {
		cfg.Server.ExternalPort = envPort
		log.Printf("Set external port from environment -> %s", cfg.Server.ExternalPort)
	} else {
		cfg.Server.ExternalPort = cfg.Server.InternalPort
		log.Printf("Environment variable PORT not set; using internal port value -> %s", cfg.Server.ExternalPort)
	}
	if err != nil {
		log.Fatalf("Error in parse config: %v", err)
	}
	return cfg
}

func InitConfig() (cfg AppConfig, err error) {

	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)
	return
}

func ParseConfig(v *viper.Viper) (*AppConfig, error) {
	var cfg AppConfig
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
	}
	return &cfg, nil

}

func LoadConfig(filename, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found.")
		}

		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {

	return "internal/config/config-development"
}
