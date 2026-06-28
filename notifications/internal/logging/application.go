package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/notifications/internal/application"
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

func (a Application) NotifyOrderCreated(ctx context.Context, notify application.OrderCreated) (err error) {
	a.logger.Info(logger.Notifications, logger.NotifyOrderCreated, "--> Notifications.NotifyOrderCreated", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Notifications, logger.NotifyOrderCreated, errors.Wrap(err, "<-- Notifications.NotifyOrderCreated").Error(), nil)
			return
		}
		a.logger.Info(logger.Notifications, logger.NotifyOrderCreated, "<-- Notifications.NotifyOrderCreated", nil)
	}()
	return a.App.NotifyOrderCreated(ctx, notify)
}

func (a Application) NotifyOrderCanceled(ctx context.Context, notify application.OrderCanceled) (err error) {
	a.logger.Info(logger.Notifications, logger.NotifyOrderCanceled, "--> Notifications.NotifyOrderCanceled", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Notifications, logger.NotifyOrderCanceled, errors.Wrap(err, "<-- Notifications.NotifyOrderCanceled").Error(), nil)
			return
		}
		a.logger.Info(logger.Notifications, logger.NotifyOrderCanceled, "<-- Notifications.NotifyOrderCanceled", nil)
	}()
	return a.App.NotifyOrderCanceled(ctx, notify)
}

func (a Application) NotifyOrderReady(ctx context.Context, notify application.OrderReady) (err error) {
	a.logger.Info(logger.Notifications, logger.NotifyOrderReady, "--> Notifications.NotifyOrderReady", nil)
	defer func() {
		if err != nil {
			a.logger.Info(logger.Notifications, logger.NotifyOrderReady, errors.Wrap(err, "<-- Notifications.NotifyOrderReady").Error(), nil)
			return
		}
		a.logger.Info(logger.Notifications, logger.NotifyOrderReady, "<-- Notifications.NotifyOrderReady", nil)
	}()

	return a.App.NotifyOrderReady(ctx, notify)
}
