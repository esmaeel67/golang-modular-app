package logging

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
)

type EventHandlers[T ddd.Event] struct {
	ddd.EventHandler[T]
	label  string
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*EventHandlers[ddd.Event])(nil)

func LogDomainEventHandlerAccess[T ddd.Event](handlers ddd.EventHandler[T], label string, logger logger.Logger) EventHandlers[T] {
	return EventHandlers[T]{
		EventHandler: handlers,
		label:        label,
		logger:       logger,
	}
}

func (h EventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	messageIn := fmt.Sprintf("--> Ordering.%s.On(%s)", h.label, event.EventName())
	h.logger.Info(logger.Orders, logger.OrderHandleEvent, messageIn, nil)
	defer func() {
		messageOut := fmt.Sprintf("<-- Ordering.%s.On(%s)", h.label, event.EventName())
		h.logger.Info(logger.Orders, logger.OrderHandleEvent, messageOut, nil)
	}()
	return h.EventHandler.HandleEvent(ctx, event)
}
