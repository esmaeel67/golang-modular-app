package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/stores/internal/application"
	"github.com/esmaeel67/golang-modular-app/stores/internal/application/commands"
	"github.com/esmaeel67/golang-modular-app/stores/internal/application/queries"
	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
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

func (a Application) CreateStore(ctx context.Context, cmd commands.CreateStore) (err error) {
	a.logger.Info(logger.Stores, logger.CreateStore, "--> Stores.CreateStore", nil)
	defer func() {
		a.logger.Info(logger.Stores, logger.CreateStore, "-<-- Stores.CreateStore", nil)
	}()
	return a.App.CreateStore(ctx, cmd)
}

func (a Application) GetStore(ctx context.Context, query queries.GetStore) (store *domain.Store, err error) {
	a.logger.Info(logger.Stores, logger.GetStore, "--> Stores.GetStore", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.GetStore, errors.Wrap(err, "<-- Stores.GetStore").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.GetStore, "<-- Stores.GetStore", nil)
	}()
	return a.App.GetStore(ctx, query)
}

func (a Application) GetStores(ctx context.Context, query queries.GetStores) (stores []*domain.Store, err error) {
	a.logger.Info(logger.Stores, logger.GetStores, "-> Stores.GetStores", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.GetStores, errors.Wrap(err, "<-- Stores.GetStores").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.GetStore, "<-- Stores.GetStores", nil)
	}()
	return a.App.GetStores(ctx, query)

}
