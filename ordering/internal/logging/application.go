package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application/commands"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application/queries"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
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
	a.logger.Info(logger.Orders, logger.CreateOrder, "--> Ordering.CreateOrder", nil)
	defer func() {
		a.logger.Info(logger.Orders, logger.CreateOrder, "<-- Ordering.CreateOrder", nil)
	}()

	return a.App.CreateOrder(ctx, cmd)
}

func (a Application) CancelOrder(ctx context.Context, cmd commands.CancelOrderCommand) (err error) {
	a.logger.Info(logger.Orders, logger.CancelOrder, "--> Ordering.CancelOrder", nil)
	defer func() {
		a.logger.Info(logger.Orders, logger.CancelOrder, "<-- Ordering.CancelOrder", nil)
	}()

	return a.App.CancelOrder(ctx, cmd)
}

func (a Application) ReadyOrder(ctx context.Context, cmd commands.ReadyOrderCommand) (err error) {
	a.logger.Info(logger.Orders, logger.ReadyOrder, "--> Ordering.ReadyOrder", nil)
	defer func() {
		a.logger.Info(logger.Orders, logger.ReadyOrder, "<-- Ordering.ReadyOrder", nil)
	}()
	return a.App.ReadyOrder(ctx, cmd)
}

func (a Application) CompleteOrder(ctx context.Context, cmd commands.CompleteOrderCommand) (err error) {
	a.logger.Info(logger.Orders, logger.CompleteOrder, "--> Ordering.CompleteOrder", nil)
	defer func() {
		a.logger.Info(logger.Orders, logger.CompleteOrder, "<-- Ordering.CompleteOrder", nil)
	}()
	return a.App.CompleteOrder(ctx, cmd)
}

func (a Application) GetOrder(ctx context.Context, query queries.GetOrderQuery) (order *domain.Order, err error) {
	a.logger.Info(logger.Orders, logger.GetOrder, "--> Ordering.GetOrder", nil)
	defer func() {
		a.logger.Info(logger.Orders, logger.CancelOrder, "<-- Ordering.GetOrder", nil)
	}()
	return a.App.GetOrder(ctx, query)
}
