package application

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/ordering/internal/application/commands"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
		CreateOrder(ctx context.Context, cmd commands.CreateOrderCommand) error
	}
	Queries interface {
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateOrderHandler
	}
	appQueries struct {
	}
)

var _ App = (*Application)(nil)

func New(orders domain.OrderRepository, customers domain.CustomerRepository, payments domain.PaymentRepository,
	invoices domain.InvoiceRepository, shopping domain.ShoppingRepository, notifications domain.NotificationRepository) *Application {
	return &Application{
		appCommands: appCommands{
			CreateOrderHandler: commands.NewCreateOrderHandler(orders, customers, payments, shopping, notifications),
		},
		appQueries: appQueries{},
	}
}
