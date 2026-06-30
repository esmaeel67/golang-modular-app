package logging

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/baskets/internal/application"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
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

func (h DomainEventHandlers) OnBasketStarted(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Baskets, logger.OnBasketStarted, "--> Baskets.OnBasketStarted", nil)
	defer func() {
		h.logger.Info(logger.Baskets, logger.OnBasketStarted, "<-- Baskets.OnBasketStarted", nil)
	}()

	return h.DomainEventHandlers.OnBasketStarted(ctx, event)
}

func (h DomainEventHandlers) OnBasketItemAdded(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Baskets, logger.OnBasketItemAdded, "--> Baskets.OnBasketItemAdded", nil)
	defer func() {
		h.logger.Info(logger.Baskets, logger.OnBasketItemAdded, "<-- Baskets.OnBasketItemAdded", nil)
	}()

	return h.DomainEventHandlers.OnBasketItemAdded(ctx, event)
}

func (h DomainEventHandlers) OnBasketItemRemoved(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Baskets, logger.OnBasketItemRemoved, "--> Baskets.OnBasketItemRemoved", nil)
	defer func() {
		h.logger.Info(logger.Baskets, logger.OnBasketItemRemoved, "<-- Baskets.OnBasketItemRemoved", nil)
	}()

	return h.DomainEventHandlers.OnBasketItemRemoved(ctx, event)
}

func (h DomainEventHandlers) OnBasketCanceled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Baskets, logger.OnBasketCanceled, "--> Baskets.OnBasketCanceled", nil)
	defer func() {
		h.logger.Info(logger.Baskets, logger.OnBasketCanceled, "<-- Baskets.OnBasketCanceled", nil)
	}()

	return h.DomainEventHandlers.OnBasketCanceled(ctx, event)
}

func (h DomainEventHandlers) OnBasketCheckedOut(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info(logger.Baskets, logger.OnBasketCheckedOut, "--> Baskets.OnBasketCheckedOut", nil)
	defer func() {
		h.logger.Info(logger.Baskets, logger.OnBasketCheckedOut, "<-- Baskets.OnBasketCheckedOut", nil)
	}()
	return h.DomainEventHandlers.OnBasketCheckedOut(ctx, event)
}
