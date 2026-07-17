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

func LogEventHandlerAccess[T ddd.Event](handlers ddd.EventHandler[T], label string, logger logger.Logger) EventHandlers[T] {
	return EventHandlers[T]{
		EventHandler: handlers,
		label:        label,
		logger:       logger,
	}
}

func (h EventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	messageIn := fmt.Sprintf("--> Search.%s.On(%s)", h.label, event.EventName())
	h.logger.Info(logger.Search, logger.HandleEvent, messageIn, nil)
	defer func() {
		messageOut := fmt.Sprintf("<-- Search.%s.On(%s)", h.label, event.EventName())
		h.logger.Info(logger.Search, logger.HandleEvent, messageOut, nil)
	}()

	return h.EventHandler.HandleEvent(ctx, event)
}
