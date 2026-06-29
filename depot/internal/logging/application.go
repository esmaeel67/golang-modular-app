package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/internal/application"
	"github.com/esmaeel67/golang-modular-app/depot/internal/application/commands"
	"github.com/esmaeel67/golang-modular-app/depot/internal/application/queries"
	"github.com/esmaeel67/golang-modular-app/depot/internal/domain"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/stackus/errors"
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

func (a Application) CreateShoppingList(ctx context.Context, cmd commands.CreateShoppingListCommand) (err error) {
	a.logger.Info(logger.Depot, logger.CreateShoppingList, "--> Depot.CreateShoppingList", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Depot, logger.CreateShoppingList, errors.Wrap(err, "<-- Depot.CreateShoppingList").Error(), nil)
			return
		}
		a.logger.Info(logger.Depot, logger.CreateShoppingList, "<-- Depot.CreateShoppingList", nil)
	}()
	return a.App.CreateShoppingList(ctx, cmd)
}

func (a Application) CancelShoppingList(ctx context.Context, cmd commands.CancelShoppingListCommand) (err error) {
	a.logger.Info(logger.Depot, logger.CancelShoppingList, "--> Depot.CancelShoppingList", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Depot, logger.CancelShoppingList, errors.Wrap(err, "<-- Depot.CancelShoppingList").Error(), nil)
			return
		}
		a.logger.Info(logger.Depot, logger.CancelShoppingList, "<-- Depot.CancelShoppingList", nil)
	}()
	return a.App.CancelShoppingList(ctx, cmd)
}

func (a Application) AssignShoppingList(ctx context.Context, cmd commands.AssignShoppingListCommand) (err error) {
	a.logger.Info(logger.Depot, logger.CancelShoppingList, "--> Depot.AssignShoppingList", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Depot, logger.CancelShoppingList, errors.Wrap(err, "<-- Depot.AssignShoppingList").Error(), nil)
			return
		}
		a.logger.Info(logger.Depot, logger.CancelShoppingList, "<-- Depot.AssignShoppingList", nil)
	}()
	return a.App.AssignShoppingList(ctx, cmd)
}

func (a Application) CompleteShoppingList(ctx context.Context, cmd commands.CompleteShoppingListCommand) (err error) {
	a.logger.Info(logger.Depot, logger.CompleteShoppingList, "--> Depot.CompleteShoppingList", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Depot, logger.CompleteShoppingList, errors.Wrap(err, "<-- Depot.CompleteShoppingList").Error(), nil)
			return
		}
		a.logger.Info(logger.Depot, logger.CompleteShoppingList, "<-- Depot.CompleteShoppingList", nil)
	}()
	return a.App.CompleteShoppingList(ctx, cmd)
}

func (a Application) GetShoppingList(ctx context.Context, query queries.GetShoppingList) (list *domain.ShoppingList,
	err error,
) {
	a.logger.Info(logger.Depot, logger.GetShoppingList, "--> Depot.GetShoppingList", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Depot, logger.GetShoppingList, errors.Wrap(err, "<-- Depot.GetShoppingList").Error(), nil)
			return
		}
		a.logger.Info(logger.Depot, logger.GetShoppingList, "<-- Depot.GetShoppingList", nil)
	}()

	return a.App.GetShoppingList(ctx, query)
}
