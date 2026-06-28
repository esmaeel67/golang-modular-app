package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/customers/internal/application"
	"github.com/esmaeel67/golang-modular-app/customers/internal/domain"
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

func (a Application) RegisterCustomer(ctx context.Context, register application.RegisterCustomer) (err error) {
	a.logger.Info(logger.Customer, logger.RegisterCustomer, "--> Customers.RegisterCustomer", nil)
	defer func() {
		a.logger.Info(logger.Customer, logger.RegisterCustomer, "<-- Customers.RegisterCustomer", nil)
	}()

	return a.App.RegisterCustomer(ctx, register)
}

func (a Application) AuthorizeCustomer(ctx context.Context, authorize application.AuthorizeCustomer) (err error) {
	a.logger.Info(logger.Customer, logger.AuthorizeCustomer, "--> Customers.AuthorizeCustomer", nil)
	defer func() {
		a.logger.Info(logger.Customer, logger.AuthorizeCustomer, "<-- Customers.AuthorizeCustomer", nil)
	}()

	return a.App.AuthorizeCustomer(ctx, authorize)
}

func (a Application) GetCustomer(ctx context.Context, get application.GetCustomer) (customer *domain.Customer, err error) {
	a.logger.Info(logger.Customer, logger.GetCustomer, "--> Customers.GetCustomer", nil)
	defer func() {
		a.logger.Info(logger.Customer, logger.GetCustomer, "<-- Customers.GetCustomer", nil)
	}()

	return a.App.GetCustomer(ctx, get)
}

func (a Application) EnableCustomer(ctx context.Context, enable application.EnableCustomer) (err error) {
	a.logger.Info(logger.Customer, logger.EnableCustomer, "--> Customers.EnableCustomer", nil)
	defer func() {
		a.logger.Info(logger.Customer, logger.EnableCustomer, "<-- Customers.EnableCustomer", nil)
	}()

	return a.App.EnableCustomer(ctx, enable)
}

func (a Application) DisableCustomer(ctx context.Context, disable application.DisableCustomer) (err error) {
	a.logger.Info(logger.Customer, logger.DisableCustomer, "--> Customers.DisableCustomer", nil)
	defer func() {
		a.logger.Info(logger.Customer, logger.DisableCustomer, "<-- Customers.DisableCustomer", nil)
	}()

	return a.App.DisableCustomer(ctx, disable)
}
