package logging

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
)

type CommandHandlers[T ddd.Command] struct {
	ddd.CommandHandler[T]
	label  string
	logger logger.Logger
}

func LogCommandHandlerAccess[T ddd.Command](handlers ddd.CommandHandler[T], label string, logger logger.Logger) ddd.CommandHandler[T] {
	return CommandHandlers[T]{
		CommandHandler: handlers,
		label:          label,
		logger:         logger,
	}
}

func (h CommandHandlers[T]) HandleCommand(ctx context.Context, command T) (reply ddd.Reply, err error) {
	messageIn := fmt.Sprintf("--> Depot.%s.On(%s)", h.label, command.CommandName())
	h.logger.Info(logger.Depot, logger.CreateShoppingList, messageIn, nil)
	defer func() {
		messageOut := fmt.Sprintf("<-- Depot.%s.On(%s)", h.label, command.CommandName())
		h.logger.Info(logger.Depot, logger.CreateShoppingList, messageOut, nil)
	}()
	return h.CommandHandler.HandleCommand(ctx, command)
}
