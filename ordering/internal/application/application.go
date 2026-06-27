package application

import "github.com/esmaeel67/golang-modular-app/ordering/internal/domain"

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
	}
	Queries interface {
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
	}
	appQueries struct {
	}
)

var _ App = (*Application)(nil)

func New(orders domain.OrderRepository, customers domain.CustomerRepository, payments domain.PaymentRepository,
	invoices domain.InvoiceRepository, shopping domain.ShoppingRepository) *Application {
	return &Application{
		appCommands: appCommands{},
		appQueries:  appQueries{},
	}
}
