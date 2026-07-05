package ordering

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/es"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	pg "github.com/esmaeel67/golang-modular-app/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/registry"
	"github.com/esmaeel67/golang-modular-app/internal/registry/serdes"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/handlers"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/logging"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	reg := registry.New()
	err = registrations(reg)
	if err != nil {
		return err
	}

	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("snapshots", mono.DB(), reg),
	)
	orders := es.NewAggregateRepository[*domain.Order](domain.OrderAggregate, reg, aggregateStore)
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	shopping := grpc.NewShoppingRepository(conn)
	notifications := grpc.NewNotificationRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, invoices, shopping)
	app = logging.NewApplication(app, mono.Logger())

	// setup application handlers
	notificationHandlers := logging.LogDomainEventHandlerAccess(
		application.NewNotificationHandlers(notifications),
		"Notification",
		mono.Logger(),
	)
	invoiceHandlers := logging.LogDomainEventHandlerAccess(
		application.NewInvoiceHandlers(invoices),
		"Invoice",
		mono.Logger(),
	)
	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	//TODO: http

	//TODO: swagger

	handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)
	handlers.RegisterInvoiceHandlers(invoiceHandlers, domainDispatcher)
	return nil
}

func registrations(reg registry.Registry) error {
	serde := serdes.NewJsonSerde(reg)

	// Order
	if err := serde.Register(domain.Order{}, func(v interface{}) error {
		order := v.(*domain.Order)
		order.Aggregate = es.NewAggregate("", domain.OrderAggregate)
		return nil
	}); err != nil {
		return err
	}

	// order event
	if err := serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderReadied{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCompleted{}); err != nil {
		return err
	}

	// order snapshot
	if err := serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return err
	}
	return nil
}
