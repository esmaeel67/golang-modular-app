package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/payments/internal/application"
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

func (a Application) AuthorizePayment(ctx context.Context, authorize application.AuthorizePayment) (err error) {
	a.logger.Info(logger.Payments, logger.AuthorizePayment, "--> Payments.AuthorizePayment", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Payments, logger.AuthorizePayment, errors.Wrap(err, "<-- Payments.AuthorizePayment").Error(), nil)
			return
		}
		a.logger.Info(logger.Payments, logger.AuthorizePayment, "<-- Payments.AuthorizePayment", nil)
	}()
	return a.App.AuthorizePayment(ctx, authorize)
}

func (a Application) ConfirmPayment(ctx context.Context, confirm application.ConfirmPayment) (err error) {
	a.logger.Info(logger.Payments, logger.ConfirmPayment, "--> Payments.ConfirmPayment", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Payments, logger.ConfirmPayment, errors.Wrap(err, "<-- Payments.ConfirmPayment").Error(), nil)
			return
		}
		a.logger.Info(logger.Payments, logger.ConfirmPayment, "<-- Payments.ConfirmPayment", nil)
	}()
	return a.App.ConfirmPayment(ctx, confirm)
}

func (a Application) CreateInvoice(ctx context.Context, create application.CreateInvoice) (err error) {
	a.logger.Info(logger.Payments, logger.CreateInvoice, "--> Payments.CreateInvoice", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Payments, logger.CreateInvoice, errors.Wrap(err, "<-- Payments.CreateInvoice").Error(), nil)
			return
		}
		a.logger.Info(logger.Payments, logger.CreateInvoice, "<-- Payments.CreateInvoice", nil)
	}()
	return a.App.CreateInvoice(ctx, create)
}

func (a Application) AdjustInvoice(ctx context.Context, adjust application.AdjustInvoice) (err error) {
	a.logger.Info(logger.Payments, logger.AdjustInvoice, "--> Payments.AdjustInvoice", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Payments, logger.AdjustInvoice, errors.Wrap(err, "<-- Payments.AdjustInvoice").Error(), nil)
			return
		}
		a.logger.Info(logger.Payments, logger.AdjustInvoice, "<-- Payments.AdjustInvoice", nil)
	}()
	return a.App.AdjustInvoice(ctx, adjust)
}

func (a Application) PayInvoice(ctx context.Context, pay application.PayInvoice) (err error) {
	a.logger.Info(logger.Payments, logger.PayInvoice, "--> Payments.PayInvoice", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Payments, logger.PayInvoice, errors.Wrap(err, "<-- Payments.PayInvoice").Error(), nil)
			return
		}
		a.logger.Info(logger.Payments, logger.PayInvoice, "<-- Payments.PayInvoice", nil)
	}()
	return a.App.PayInvoice(ctx, pay)
}

func (a Application) CancelInvoice(ctx context.Context, cancel application.CancelInvoice) (err error) {
	a.logger.Info(logger.Payments, logger.CancelInvoice, "--> Payments.CancelInvoice", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Payments, logger.CancelInvoice, errors.Wrap(err, "<-- Payments.CancelInvoice").Error(), nil)
			return
		}
		a.logger.Info(logger.Payments, logger.CancelInvoice, "<-- Payments.CancelInvoice", nil)
	}()
	return a.App.CancelInvoice(ctx, cancel)
}
