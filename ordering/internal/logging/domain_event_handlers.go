package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application"
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

func (h DomainEventHandlers) OnOrderCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Orders, logger.OnOrderCreated, "--> Ordering.OnOrderCreated", nil)
	defer func() {
		h.logger.Info(logger.Orders, logger.OnOrderCreated, "<-- Ordering.OnOrderCreated", nil)
	}()
	return h.DomainEventHandlers.OnOrderCreated(ctx, event)
}

func (h DomainEventHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Orders, logger.OnOrderReadied, "--> Ordering.OnOrderReadied", nil)
	defer func() {
		h.logger.Info(logger.Orders, logger.OnOrderReadied, "<-- Ordering.OnOrderReadied", nil)
	}()
	return h.DomainEventHandlers.OnOrderReadied(ctx, event)
}

func (h DomainEventHandlers) OnOrderCanceled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Orders, logger.OnOrderCanceled, "--> Ordering.OnOrderCanceled", nil)
	defer func() {
		h.logger.Info(logger.Orders, logger.OnOrderCanceled, "<-- Ordering.OnOrderCanceled", nil)
	}()
	return h.DomainEventHandlers.OnOrderCanceled(ctx, event)
}
