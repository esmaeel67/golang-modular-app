package notifications

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/customers/customerspb"
	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/jetstream"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/internal/registry"
	"github.com/esmaeel67/golang-modular-app/notifications/internal/application"
	"github.com/esmaeel67/golang-modular-app/notifications/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/notifications/internal/handlers"
	"github.com/esmaeel67/golang-modular-app/notifications/internal/logging"
	"github.com/esmaeel67/golang-modular-app/notifications/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/ordering/orderingpb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {

	// setup Driven adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}

	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()))
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := postgres.NewCustomerCacheRepository("customers_cache", mono.DB(), grpc.NewCustomerRepository(conn))

	// setup application

	app := logging.LogApplicationAccess(application.New(customers), mono.Logger())

	integrationEventHandlers := logging.LogEventHandlerAccess(
		handlers.NewIntegrationEventHandlers(app, customers),
		"IntegrationEvents", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	return nil
}
