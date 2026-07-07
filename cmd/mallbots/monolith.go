package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/esmaeel67/golang-modular-app/internal/config"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/internal/waiter"
	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type app struct {
	cfg     config.AppConfig
	db      *sql.DB
	nc      *nats.Conn
	js      nats.JetStreamContext
	logger  logger.Logger
	modules []monolith.Module
	mux     *chi.Mux
	rpc     *grpc.Server
	waiter  waiter.Waiter
}

func (a *app) Config() config.AppConfig {
	return a.cfg
}
func (a *app) DB() *sql.DB {
	return a.db
}

func (a *app) JS() nats.JetStreamContext {
	return a.js
}

func (a *app) Logger() logger.Logger {
	return a.logger
}

func (a *app) Mux() *chi.Mux {
	return a.mux
}

func (a *app) RPC() *grpc.Server {
	return a.rpc
}
func (a *app) Waiter() waiter.Waiter {
	return a.waiter
}

func (a *app) startupModules() error {

	for _, module := range a.modules {
		if err := module.Startup(a.Waiter().Context(), a); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) waitForWeb(ctx context.Context) error {

	webServer := &http.Server{
		Addr:    a.cfg.Web.Address(),
		Handler: a.mux,
	}

	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		fmt.Println("Web server started")
		defer fmt.Println("Web server shutdown")
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("Web server to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.ShutdownTimeout)
		defer cancel()

		if err := webServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})
	return group.Wait()
}

func (a *app) waitForRPC(ctx context.Context) error {
	fmt.Println("a.cfg.Rpc.Address() : ", a.cfg.Rpc.Address())
	listener, err := net.Listen("tcp", a.cfg.Rpc.Address())
	if err != nil {
		return err
	}
	//
	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		fmt.Println("Rpc server started")
		defer fmt.Println("Rpc server shutdown")

		if err := a.RPC().Serve(listener); err != nil && err != grpc.ErrServerStopped {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("Rpc server to be shutdown")
		stopped := make(chan struct{})
		go func() {
			a.RPC().GracefulStop()
			close(stopped)
		}()

		timeout := time.NewTimer(a.cfg.ShutdownTimeout)
		defer timeout.Stop()

		select {
		case <-timeout.C:
			a.RPC().Stop()
			return fmt.Errorf("rpc server failed to stop gracefully")
		case <-stopped:
			return nil
		}
	})

	return group.Wait()
}

func (a *app) waitForStream(ctx context.Context) error {
	closed := make(chan struct{})
	a.nc.SetClosedHandler(func(c *nats.Conn) {
		close(closed)
	})

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("Message stream started")
		defer fmt.Println("message stream stopped")
		<-closed
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		return a.nc.Drain()
	})
	return group.Wait()
}
