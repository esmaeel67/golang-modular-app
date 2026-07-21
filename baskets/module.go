package baskets

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/baskets/basketspb"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/application"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/domain"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/handlers"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/logging"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/es"
	"github.com/esmaeel67/golang-modular-app/internal/jetstream"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	pg "github.com/esmaeel67/golang-modular-app/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/registry"
	"github.com/esmaeel67/golang-modular-app/internal/registry/serdes"
	"github.com/esmaeel67/golang-modular-app/stores/storespb"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	err = registrations(reg)
	if err != nil {
		return err
	}

	if err = basketspb.Registrations(reg); err != nil {
		return err
	}

	if err = storespb.Registrations(reg); err != nil {
		return err
	}

	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("events", mono.DB(), reg),
		pg.NewSnapshotStore("snapshots", mono.DB(), reg),
	)
	baskets := es.NewAggregateRepository[*domain.Basket](domain.BasketAggregate, reg, aggregateStore)

	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	stores := postgres.NewStoreCacheRepository("store_cache", mono.DB(), grpc.NewStoreRepository(conn))
	products := postgres.NewProductCacheRepository("products_cache", mono.DB(), grpc.NewProductRepository(conn))

	// setup application
	app := logging.LogApplicationAccess(
		application.New(baskets, stores, products, domainDispatcher),
		mono.Logger(),
	)
	domainEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewDomainEventHandlers(eventStream),
		"DomainEvents", mono.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewIntegrationEventHandlers(stores, products),
		"IntegrationEvents", mono.Logger(),
	)
	// orderHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
	// 	application.NewOrderHandler(orders),
	// 	"Order", mono.Logger(),
	// )
	// storeHandlers := logging.LogEventHandlerAccess[ddd.Event](
	// 	application.NewStoreHandlers(mono.Logger()),
	// 	"Store", mono.Logger(),
	// )
	// productHandlers := logging.LogEventHandlerAccess[ddd.Event](
	// 	application.NewProductHandlers(mono.Logger()),
	// 	"Product", mono.Logger(),
	// )

	// Print the type of registrar
	fmt.Printf("🔍 Registrar type: %T\n", mono.RPC())
	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	// handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)
	// if err = handlers.RegisterStoreHandlers(storeHandlers, eventStream); err != nil {
	// 	return err
	// }
	// if err = handlers.RegisterProductHandlers(productHandlers, eventStream); err != nil {
	// 	return err
	// }

	return nil
}

func registrations(reg registry.Registry) error {
	serde := serdes.NewJsonSerde(reg)

	// basket
	if err := serde.Register(domain.Basket{}, func(v interface{}) error {
		basket := v.(*domain.Basket)
		basket.Items = make(map[string]domain.Item)
		return nil
	}); err != nil {
		return err
	}

	// basket events
	if err := serde.Register(domain.BasketStarted{}); err != nil {
		return err
	}
	if err := serde.Register(domain.BasketCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(domain.BasketCheckedOut{}); err != nil {
		return err
	}
	if err := serde.Register(domain.BasketItemAdded{}); err != nil {
		return err
	}
	if err := serde.Register(domain.BasketItemRemoved{}); err != nil {
		return err
	}
	// basket snapshots
	if err := serde.RegisterKey(domain.BasketV1{}.SnapshotName(), domain.BasketV1{}); err != nil {
		return err
	}

	return nil
}
