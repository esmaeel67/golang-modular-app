package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/esmaeel67/golang-modular-app/baskets"
	"github.com/esmaeel67/golang-modular-app/customers"
	"github.com/esmaeel67/golang-modular-app/depot"
	"github.com/esmaeel67/golang-modular-app/internal/config"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/internal/rpc"
	"github.com/esmaeel67/golang-modular-app/internal/waiter"
	"github.com/esmaeel67/golang-modular-app/internal/web"
	"github.com/esmaeel67/golang-modular-app/migrations"
	"github.com/esmaeel67/golang-modular-app/notifications"
	"github.com/esmaeel67/golang-modular-app/ordering"
	"github.com/esmaeel67/golang-modular-app/payments"
	"github.com/esmaeel67/golang-modular-app/stores"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nats-io/nats.go"
	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("startUp", err.Error())
		os.Exit(1)
	}
}

func run() error {

	cfg := config.GetConfig()

	m := app{cfg: *cfg}
	// setup database connection // cfg.PG.Conn
	var err error
	m.db, err = sql.Open("pgx", cfg.Postgres.GetPostgresConn())
	// m.db, err = sql.Open("pgx", "postgres://postgres:admin@localhost:5432/mallbots?sslmode=disable")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.db)
	// migration database
	err = migrateDB(m.db, *cfg)
	if err != nil {
		return err
	}
	fmt.Println("FFFFFFFFFFFFF :", cfg.Nats.GetUrl())
	// init nats & jetstream
	m.nc, err = nats.Connect(cfg.Nats.GetUrl())
	if err != nil {
		return err
	}
	defer m.nc.Close()
	m.js, err = initJetStream(cfg.Nats, m.nc)
	if err != nil {
		return err
	}

	// logger config
	m.logger = logger.NewLogger(cfg)

	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	m.modules = []monolith.Module{
		&baskets.Module{},
		&customers.Module{},
		&depot.Module{},
		&notifications.Module{},
		&stores.Module{},
		&ordering.Module{},
		&payments.Module{},
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	m.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started app")
	defer fmt.Println("stopped app")

	m.waiter.Add(m.waitForWeb, m.waitForRPC, m.waitForStream)

	return m.waiter.Wait()
}

func migrateDB(db *sql.DB, cfg config.AppConfig) error {
	goose.SetVerbose(cfg.Goose.Debug)
	goose.SetBaseFS(migrations.FS)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "."); err != nil {
		return err
	}

	return nil
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)
	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}

func initJetStream(cfg config.NatsConfig, nc *nats.Conn) (nats.JetStreamContext, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     cfg.Stream,
		Subjects: []string{fmt.Sprintf("%s.>", cfg.Stream)},
	})
	return js, err
}
