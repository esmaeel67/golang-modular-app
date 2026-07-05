package customers

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/customers/internal/application"
	"github.com/esmaeel67/golang-modular-app/customers/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/customers/internal/logging"
	"github.com/esmaeel67/golang-modular-app/customers/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	customers := postgres.NewCustomerRepository("customers", mono.DB())

	var app application.App

	app = application.New(customers, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	//TODO: register rest api

	//TODO: register swagger server api and config

	return nil
}
