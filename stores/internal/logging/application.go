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

func (a Application) EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) (err error) {
	a.logger.Info(logger.Stores, logger.EnableParticipation, "--> Stores.EnableParticipation", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.EnableParticipation, errors.Wrap(err, "<-- Stores.EnableParticipation").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.EnableParticipation, "<-- Stores.EnableParticipation", nil)
	}()

	return a.App.EnableParticipation(ctx, cmd)
}

func (a Application) DisableParticipation(ctx context.Context, cmd commands.DisableParticipation) (err error) {
	a.logger.Info(logger.Stores, logger.DisableParticipation, "--> Stores.DisableParticipation", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.EnableParticipation, errors.Wrap(err, "<-- Stores.DisableParticipation").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.EnableParticipation, "<-- Stores.DisableParticipation", nil)
	}()
	return a.App.DisableParticipation(ctx, cmd)
}

func (a Application) AddProduct(ctx context.Context, cmd commands.AddProduct) (err error) {
	a.logger.Info(logger.Stores, logger.AddProduct, "--> Stores.AddProduct", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.GetParticipatingStores, errors.Wrap(err, "<-- Stores.AddProduct").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.GetParticipatingStores, "<-- Stores.AddProduct", nil)
	}()
	return a.App.AddProduct(ctx, cmd)
}

func (a Application) RemoveProduct(ctx context.Context, cmd commands.RemoveProductCommand) (err error) {
	a.logger.Info(logger.Stores, logger.AddProduct, "--> Stores.RemoveProduct", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.GetParticipatingStores, errors.Wrap(err, "<-- Stores.RemoveProduct").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.GetParticipatingStores, "<-- Stores.RemoveProduct", nil)
	}()
	return a.App.RemoveProduct(ctx, cmd)
}

func (a Application) GetParticipatingStores(ctx context.Context, query queries.GetParticipatingStores) (store []*domain.Store, err error) {
	a.logger.Info(logger.Stores, logger.GetParticipatingStores, "--> Stores.GetParticipatingStores", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.GetParticipatingStores, errors.Wrap(err, "<-- Stores.GetParticipatingStores").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.GetParticipatingStores, "<-- Stores.GetParticipatingStores", nil)
	}()

	return a.App.GetParticipatingStores(ctx, query)
}

func (a Application) GetCatalog(ctx context.Context, query queries.GetCatalogQuery) (products []*domain.Product, err error) {
	a.logger.Info(logger.Stores, logger.GetCatalog, "--> Stores.GetCatalog", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.GetCatalog, errors.Wrap(err, "<-- Stores.GetCatalog").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.GetCatalog, "<-- Stores.GetCatalog", nil)
	}()
	return a.App.GetCatalog(ctx, query)
}

func (a Application) GetProduct(ctx context.Context, query queries.GetProductQuery) (product *domain.Product, err error) {
	a.logger.Info(logger.Stores, logger.GetProduct, "--> Stores.GetProduct", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Stores, logger.GetProduct, errors.Wrap(err, "<-- Stores.GetParticipatingStores").Error(), nil)
			return
		}
		a.logger.Info(logger.Stores, logger.GetProduct, "<-- Stores.GetParticipatingStores", nil)
	}()
	return a.App.GetProduct(ctx, query)
}
