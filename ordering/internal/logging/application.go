package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application/commands"
)

type Application struct {
	application.App
	logger logger.Logger
}

var _ application.App = (*Application)(nil)

func NewApplication(application application.App, logger logger.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) CreateOrder(ctx context.Context, cmd commands.CreateOrderCommand) (err error) {
	a.logger.Info(logger.Customers, logger.RegisterCustomer, "--> Ordering.CreateOrder", nil)
	defer func() {
		a.logger.Info(logger.Customers, logger.RegisterCustomer, "<-- Ordering.CreateOrder", nil)
	}()

	return a.App.CreateOrder(ctx, cmd)
}
