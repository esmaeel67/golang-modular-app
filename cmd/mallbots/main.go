package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/esmaeel67/golang-modular-app/baskets"
	"github.com/esmaeel67/golang-modular-app/customers"
	"github.com/esmaeel67/golang-modular-app/internal/config"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/internal/rpc"
	"github.com/esmaeel67/golang-modular-app/internal/waiter"
	"github.com/esmaeel67/golang-modular-app/internal/web"
	"github.com/esmaeel67/golang-modular-app/stores"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
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

	var cfg config.AppConfig
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	m := app{cfg: cfg}
	// setup database connection // cfg.PG.Conn
	m.db, err = sql.Open("pgx", cfg.PG.Conn)
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

	// logger config
	m.logger = logger.NewLogger(&cfg)

	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	m.modules = []monolith.Module{
		&baskets.Module{},
		&customers.Module{},
		&stores.Module{},
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	m.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started app")
	defer fmt.Println("stopped app")

	m.waiter.Add(m.waitForWeb, m.waitForRPC)

	return m.waiter.Wait()
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)
	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
