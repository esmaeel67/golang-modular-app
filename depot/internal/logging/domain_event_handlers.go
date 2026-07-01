package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/internal/application"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/stackus/errors"
)

type DomainEventHandlers struct {
	application.DomainEventHandlers
	logger logger.Logger
}

var _ application.DomainEventHandlers = (*DomainEventHandlers)(nil)

func LogDomainEventHandlerAccess(handlers application.DomainEventHandlers, logger logger.Logger) DomainEventHandlers {
	return DomainEventHandlers{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnShoppingListCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Depot, logger.OnShoppingListCreated, "--> Depot.OnShoppingListCreated", nil)
	defer func() {
		if err != nil {
			h.logger.Info(logger.Depot, logger.OnShoppingListCreated, errors.Wrap(err, "<-- Depot.OnShoppingListCreated").Error(), nil)
			return
		}
		h.logger.Info(logger.Depot, logger.OnShoppingListCreated, "<-- Depot.OnShoppingListCreated", nil)
	}()
	return h.DomainEventHandlers.OnShoppingListCreated(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListCanceled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Depot, logger.OnShoppingListCanceled, "--> Depot.OnShoppingListCanceled", nil)
	defer func() {
		if err != nil {
			h.logger.Info(logger.Depot, logger.OnShoppingListCanceled, errors.Wrap(err, "<-- Depot.OnShoppingListCanceled").Error(), nil)
			return
		}
		h.logger.Info(logger.Depot, logger.OnShoppingListCanceled, "<-- Depot.OnShoppingListCanceled", nil)
	}()

	return h.DomainEventHandlers.OnShoppingListCanceled(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListAssigned(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Depot, logger.OnShoppingListAssigned, "--> Depot.OnShoppingListAssigned", nil)
	defer func() {
		if err != nil {
			h.logger.Info(logger.Depot, logger.OnShoppingListAssigned, errors.Wrap(err, "<-- Depot.OnShoppingListAssigned").Error(), nil)
			return
		}
		h.logger.Info(logger.Depot, logger.OnShoppingListAssigned, "<-- Depot.OnShoppingListAssigned", nil)
	}()

	return h.DomainEventHandlers.OnShoppingListAssigned(ctx, event)
}

func (h DomainEventHandlers) OnShoppingListCompleted(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Depot, logger.OnShoppingListCompleted, "--> Depot.OnShoppingListCompleted", nil)
	defer func() {
		if err != nil {
			h.logger.Info(logger.Depot, logger.OnShoppingListCompleted, errors.Wrap(err, "<-- Depot.OnShoppingListCompleted").Error(), nil)
			return
		}
		h.logger.Info(logger.Depot, logger.OnShoppingListCompleted, "<-- Depot.OnShoppingListCompleted", nil)
	}()

	return h.DomainEventHandlers.OnShoppingListCompleted(ctx, event)
}
