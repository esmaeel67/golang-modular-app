package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/baskets/internal/application"
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
