package stores

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/es"
	"github.com/esmaeel67/golang-modular-app/internal/jetstream"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	pg "github.com/esmaeel67/golang-modular-app/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/registry"
	"github.com/esmaeel67/golang-modular-app/internal/registry/serdes"
	"github.com/esmaeel67/golang-modular-app/stores/internal/application"
	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
	"github.com/esmaeel67/golang-modular-app/stores/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/stores/internal/handlers"
	"github.com/esmaeel67/golang-modular-app/stores/internal/logging"
	"github.com/esmaeel67/golang-modular-app/stores/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/stores/storespb"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	reg := registry.New()
	err := registrations(reg)
	if err != nil {
		return err
	}
	if err = storespb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(mono.Config().Nats.Stream, mono.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("snapshots", mono.DB(), reg),
	)

	stores := es.NewAggregateRepository[*domain.Store](domain.StoreAggregate, reg, aggregateStore)
	products := es.NewAggregateRepository[*domain.Product](domain.ProductAggregate, reg, aggregateStore)
	catalog := postgres.NewCatalogRepository("products", mono.DB())
	mall := postgres.NewMallRepository("stores", mono.DB())

	// setup application
	app := logging.LogApplicationAccess(
		application.New(stores, products, catalog, mall),
		mono.Logger())
	catalogHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewCatalogHandlers(catalog),
		"Catalog",
		mono.Logger(),
	)
	mallHandlers := logging.LogEventHandlerAccess(
		application.NewMallHandlers(mall),
		"Mall",
		mono.Logger(),
	)
	// added integration event handler
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", mono.Logger(),
	)

	if err := grpc.RegisterService(ctx, app, mono.RPC()); err != nil {
		return err
	}

	handlers.RegisterCatalogHandlers(catalogHandlers, domainDispatcher)
	handlers.RegisterMallHandlers(mallHandlers, domainDispatcher)
	handlers.RegisterIntegrationEventHandlers(integrationEventHandlers, domainDispatcher)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)
	// Store
	if err = serde.Register(domain.Store{}, func(v interface{}) error {
		store := v.(*domain.Store)
		store.Aggregate = es.NewAggregate("", domain.StoreAggregate)
		return nil
	}); err != nil {
		return
	}

	// store events
	if err = serde.Register(domain.StoreCreated{}); err != nil {
		return err
	}
	if err = serde.RegisterKey(domain.StoreParticipationEnabledEvent, domain.StoreParticipationToggled{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.StoreParticipationDisabledEvent, domain.StoreParticipationToggled{}); err != nil {
		return
	}
	if err = serde.Register(domain.StoreRebranded{}); err != nil {
		return
	}

	// store snapshot
	if err = serde.RegisterKey(domain.StoreV1{}.SnapshotName(), domain.StoreV1{}); err != nil {
		return
	}

	// Product
	if err = serde.Register(domain.Product{}, func(v interface{}) error {
		store := v.(*domain.Product)
		store.Aggregate = es.NewAggregate("", domain.ProductAggregate)
		return nil
	}); err != nil {
		return
	}

	// product events
	if err = serde.Register(domain.ProductAdded{}); err != nil {
		return
	}
	if err = serde.Register(domain.ProductRebranded{}); err != nil {
		return
	}

	if err = serde.RegisterKey(domain.ProductPriceIncreasedEvent, domain.ProductPriceChanged{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.ProductPriceDecreasedEvent, domain.ProductPriceChanged{}); err != nil {
		return
	}
	if err = serde.Register(domain.ProductRemoved{}); err != nil {
		return
	}
	// product snapshots
	if err = serde.RegisterKey(domain.ProductV1{}.SnapshotName(), domain.ProductV1{}); err != nil {
		return
	}

	return
}
