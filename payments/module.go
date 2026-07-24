package payments

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/jetstream"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/internal/registry"
	"github.com/esmaeel67/golang-modular-app/ordering/orderingpb"
	"github.com/esmaeel67/golang-modular-app/payments/internal/application"
	"github.com/esmaeel67/golang-modular-app/payments/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/payments/internal/handlers"
	"github.com/esmaeel67/golang-modular-app/payments/internal/logging"
	"github.com/esmaeel67/golang-modular-app/payments/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/payments/paymentspb"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {

	// set Driven adapters
	reg := registry.New()

	if err = orderingpb.Registrations(reg); err != nil {
		return err
	}
	if err = paymentspb.Registrations(reg); err != nil {
		return err
	}

	stream := jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), mono.Logger())
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	invoices := postgres.NewInvoiceRepository("invoices", mono.DB())
	payments := postgres.NewPaymentRepository("payments", mono.DB())

	// setup application

	app := logging.LogApplicationAccess(application.New(invoices, payments, domainDispatcher), mono.Logger())
	domainEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		handlers.NewDomainHandlers(eventStream),
		"DomainEvents", mono.Logger(),
	)

	integrationEventHandlers := logging.LogEventHandlerAccess(
		handlers.NewIntegrationHandlers(app),
		"IntegrationEvents", mono.Logger(),
	)

	commandHandlers := logging.LogCommandHandlerAccess[ddd.Command](
		handlers.NewCommandHandlers(app),
		"Commands", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}

	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlers(domainDispatcher, domainEventHandlers)
	if err = handlers.RegisterCommandHandlers(commandStream, commandHandlers); err != nil {
		return err
	}

	return nil
}
