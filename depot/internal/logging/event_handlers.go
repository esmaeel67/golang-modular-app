package logging

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/depot/internal/application"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/stackus/errors"
)

type EventHandlers[T ddd.Event] struct {
	application.DomainEventHandlers
	label  string
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*EventHandlers[ddd.Event])(nil)

func LogDomainEventHandlerAccess[T ddd.Event](handlers application.DomainEventHandlers, label string, logger logger.Logger) EventHandlers[T] {
	return EventHandlers[T]{
		DomainEventHandlers: handlers,
		label:               label,
		logger:              logger,
	}
}

func (h EventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	messageIn := fmt.Sprintf("--> Depot.%s.On(%s)", h.label, event.EventName())
	h.logger.Info(logger.Depot, logger.DepotHandleEvent, messageIn, nil)
	defer func() {
		messageOut := fmt.Sprintf("<-- Depot.%s.On(%s)", h.label, event.EventName())
		if err != nil {
			h.logger.Info(logger.Depot, logger.DepotHandleEvent, errors.Wrap(err, messageOut).Error(), nil)
			return
		}
		h.logger.Info(logger.Depot, logger.DepotHandleEvent, messageOut, nil)
	}()
	return h.DomainEventHandlers.OnShoppingListCreated(ctx, event)
}
