package application

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application/commands"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application/queries"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateOrder(ctx context.Context, cmd commands.CreateOrderCommand) error
		CancelOrder(ctx context.Context, cmd commands.CancelOrderCommand) error
		ReadyOrder(ctx context.Context, cmd commands.ReadyOrderCommand) error
		CompleteOrder(ctx context.Context, cmd commands.CompleteOrderCommand) error
	}
	Queries interface {
		GetOrder(ctx context.Context, query queries.GetOrderQuery) (*domain.Order, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateOrderHandler
		commands.CancelOrderHandler
		commands.ReadyOrderHandler
		commands.CompleteOrderHandler
	}
	appQueries struct {
		queries.GetOrderHandler
	}
)

var _ App = (*Application)(nil)

func New(orders domain.OrderRepository, customers domain.CustomerRepository, payments domain.PaymentRepository,
	invoices domain.InvoiceRepository, shopping domain.ShoppingRepository, domainPublisher ddd.EventPublisher) *Application {
	return &Application{
		appCommands: appCommands{
			CreateOrderHandler:   commands.NewCreateOrderHandler(orders, customers, payments, shopping, domainPublisher),
			CancelOrderHandler:   commands.NewCancelOrderHandler(orders, shopping, domainPublisher),
			ReadyOrderHandler:    commands.NewReadyOrderHandler(orders, domainPublisher),
			CompleteOrderHandler: commands.NewCompleteOrderHandler(orders, domainPublisher),
		},
		appQueries: appQueries{
			GetOrderHandler: queries.NewGetOrderHandler(orders),
		},
	}
}
