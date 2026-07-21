package search

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/customers/customerspb"
	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/jetstream"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/internal/registry"
	"github.com/esmaeel67/golang-modular-app/ordering/orderingpb"
	"github.com/esmaeel67/golang-modular-app/search/internal/application"
	"github.com/esmaeel67/golang-modular-app/search/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/search/internal/handlers"
	"github.com/esmaeel67/golang-modular-app/search/internal/logging"
	"github.com/esmaeel67/golang-modular-app/search/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/stores/storespb"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {

	reg := registry.New()
	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	if err = storespb.Registrations(reg); err != nil {
		return err
	}

	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()))
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	customers := postgres.NewCustomerCacheRepository("search_customers_cache", mono.DB(), grpc.NewCustomerRepository(conn))
	stores := postgres.NewStoreCacheRepository("search_stores_cache", mono.DB(), grpc.NewStoreRepository(conn))
	products := postgres.NewProductCacheRepository("search_products_cache", mono.DB(), grpc.NewProductRepository(conn))
	orders := postgres.NewOrderRepository("search_orders", mono.DB())

	//setup application
	app := logging.LogApplicationAccess(
		application.New(orders),
		mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(orders, customers, stores, products),
		"IntegrationEvents", mono.Logger(),
	)
	// orderHandlers := logging.LogEventHandlerAccess[ddd.Event](
	// 	application.NewOrderHandlers(orders, customers, stores, products),
	// 	"Order", mono.Logger(),
	// )
	// customerHandlers := logging.LogEventHandlerAccess[ddd.Event](
	// 	application.NewCustomerHandlers(customers),
	// 	"Customer", mono.Logger(),
	// )
	// storeHandlers := logging.LogEventHandlerAccess[ddd.Event](
	// 	application.NewStoreHandlers(stores),
	// 	"Store", mono.Logger(),
	// )
	// productHandlers := logging.LogEventHandlerAccess[ddd.Event](
	// 	application.NewProductHandlers(products),
	// 	"Product", mono.Logger(),
	// )

	// setup Driver adapters
	if err = grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}

	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}

	// if err = handlers.RegisterOrderHandlers(orderHandlers, eventStream); err != nil {
	// 	return err
	// }
	// if err = handlers.RegisterCustomerHandlers(customerHandlers, eventStream); err != nil {
	// 	return err
	// }
	// if err = handlers.RegisterStoreHandlers(storeHandlers, eventStream); err != nil {
	// 	return err
	// }
	// if err = handlers.RegisterProductHandlers(productHandlers, eventStream); err != nil {
	// 	return err
	// }

	return nil
}
