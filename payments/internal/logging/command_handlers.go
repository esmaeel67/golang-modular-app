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

var _ ddd.CommandHandler[ddd.Command] = (*CommandHandlers[ddd.Command])(nil)

func LogCommandHandlerAccess[T ddd.Command](handlers ddd.CommandHandler[T], label string, logger logger.Logger) ddd.CommandHandler[T] {
	return CommandHandlers[T]{
		CommandHandler: handlers,
		label:          label,
		logger:         logger,
	}
}

func (h CommandHandlers[T]) HandleCommand(ctx context.Context, command T) (reply ddd.Reply, err error) {
	messageIn := fmt.Sprintf("--> Payments.%s.On(%s)", h.label, command.CommandName())
	h.logger.Info(logger.Payments, logger.HandleCommand, messageIn, nil)
	defer func() {
		messageOut := fmt.Sprintf("<-- Payments.%s.On(%s)", h.label, command.CommandName())
		h.logger.Info(logger.Payments, logger.HandleCommand, messageOut, nil)
	}()
	return h.CommandHandler.HandleCommand(ctx, command)
}
