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

func LogEventHandlerAccess[T ddd.Event](handlers ddd.EventHandler[T], label string, logger logger.Logger) EventHandlers[T] {
	return EventHandlers[T]{
		EventHandler: handlers,
		label:        label,
		logger:       logger,
	}
}

func (h EventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	messageIn := fmt.Sprintf("--> Payments.%s.On(%s)", h.label, event.EventName())
	h.logger.Info(logger.Payments, logger.HandleEvent, messageIn, nil)
	defer func() {
		messageOut := fmt.Sprintf("<-- Payments.%s.On(%s)", h.label, event.EventName())
		h.logger.Info(logger.Payments, logger.HandleEvent, messageOut, nil)
	}()
	return h.EventHandler.HandleEvent(ctx, event)
}
