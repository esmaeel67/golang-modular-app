package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/search/internal/application"
	"github.com/esmaeel67/golang-modular-app/search/internal/models"
	"github.com/stackus/errors"
)

type Application struct {
	application.Application
	logger logger.Logger
}

var _ application.Application = (*Application)(nil)

func LogApplicationAccess(application application.Application, logger logger.Logger) Application {
	return Application{
		Application: application,
		logger:      logger,
	}
}

func (a Application) SearchOrders(ctx context.Context, search application.SearchOrders) (orders []*models.Order, err error) {
	a.logger.Info(logger.Search, logger.SearchOrders, "--> Search.SearchOrders", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Search, logger.SearchOrders, errors.Wrap(err, "<-- Search.SearchOrders").Error(), nil)
			return
		}
		a.logger.Info(logger.Search, logger.SearchOrders, "<-- Search.SearchOrders", nil)
	}()
	return a.Application.SearchOrders(ctx, search)
}

func (a Application) GetOrder(ctx context.Context, getOrder application.GetOrder) (order *models.Order, err error) {
	a.logger.Info(logger.Search, logger.GetOrder, "--> Search.GetOrder", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Search, logger.GetOrder, errors.Wrap(err, "<-- Search.GetOrder").Error(), nil)
			return
		}
		a.logger.Info(logger.Search, logger.GetOrder, "<-- Search.GetOrder", nil)
	}()
	return a.Application.GetOrder(ctx, getOrder)
}
