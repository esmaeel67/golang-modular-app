package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/baskets/internal/application"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/domain"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
)

type Application struct {
	application.App
	logger logger.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger logger.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}
func (a Application) StartBasket(ctx context.Context, start application.StartBasket) (err error) {
	a.logger.Info(logger.Baskets, logger.StartBasket, "--> Baskets.CancelBasket", nil)
	defer func() {
		a.logger.Info(logger.Baskets, logger.StartBasket, "<-- Baskets.StartBasket", nil)
	}()
	return a.App.StartBasket(ctx, start)
}

func (a Application) CancelBasket(ctx context.Context, cancel application.CancelBasket) (err error) {
	a.logger.Info(logger.Baskets, logger.CancelBasket, "--> Baskets.CancelBasket", nil)
	defer func() {
		a.logger.Info(logger.Baskets, logger.CancelBasket, "<-- Baskets.CancelBasket", nil)
	}()
	return a.App.CancelBasket(ctx, cancel)
}

func (a Application) CheckoutBasket(ctx context.Context, checkout application.CheckoutBasket) (err error) {
	a.logger.Info(logger.Baskets, logger.CheckoutBasket, "--> Baskets.CheckoutBasket", nil)
	defer func() {
		a.logger.Info(logger.Baskets, logger.CheckoutBasket, "<-- Baskets.CheckoutBasket", nil)
	}()
	return a.App.CheckoutBasket(ctx, checkout)
}

func (a Application) AddItem(ctx context.Context, add application.AddItem) (err error) {
	a.logger.Info(logger.Baskets, logger.BasketAddItem, "--> Baskets.AddItem", nil)
	defer func() {
		a.logger.Info(logger.Baskets, logger.BasketAddItem, "<-- Baskets.AddItem", nil)
	}()
	return a.App.AddItem(ctx, add)
}

func (a Application) RemoveItem(ctx context.Context, remove application.RemoveItem) (err error) {
	a.logger.Info(logger.Baskets, logger.BasketRemoveItem, "--> Baskets.RemoveItem", nil)
	defer func() {
		a.logger.Info(logger.Baskets, logger.BasketRemoveItem, "<-- Baskets.RemoveItem", nil)
	}()
	return a.App.RemoveItem(ctx, remove)
}

func (a Application) GetBasket(ctx context.Context, get application.GetBasket) (basket *domain.Basket, err error) {
	a.logger.Info(logger.Baskets, logger.GetBasket, "--> Baskets.GetBasket", nil)
	defer func() {
		a.logger.Info(logger.Baskets, logger.GetBasket, "<-- Baskets.GetBasket", nil)
	}()
	return a.App.GetBasket(ctx, get)
}
