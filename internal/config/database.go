package config

import (
	"fmt"
	"time"
)

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func (postgres PostgresConfig) GetPostgresConn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		postgres.Host, postgres.Port, postgres.User, postgres.Password, postgres.DbName, postgres.SSLMode)
}

type GooseConfig struct {
	Debug  bool   `default:"true"`
	Driver string `default:"postgres"`
}
