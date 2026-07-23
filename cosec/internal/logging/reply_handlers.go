package logging

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/internal/sec"
	"github.com/stackus/errors"
)

type sagaReplyHandlers[T any] struct {
	sec.Orchestrator[T]
	label  string
	logger logger.Logger
}

var _ sec.Orchestrator[any] = (*sagaReplyHandlers[any])(nil)

func LogReplyHandlerAccess[T any](orc sec.Orchestrator[T], label string, logger logger.Logger) sec.Orchestrator[T] {
	return sagaReplyHandlers[T]{
		Orchestrator: orc,
		label:        label,
		logger:       logger,
	}
}

func (h sagaReplyHandlers[T]) HandleReply(ctx context.Context, reply ddd.Reply) (err error) {
	messageIn := fmt.Sprintf("--> COSEC.%s.On(%s)", h.label, reply.ReplyName())
	h.logger.Info(logger.COSEC, logger.HandleReply, messageIn, nil)
	defer func() {
		messageOut := fmt.Sprintf("<-- COSEC.%s.On(%s)", h.label, reply.ReplyName())
		if err != nil {
			h.logger.Info(logger.COSEC, logger.HandleReply, errors.Wrap(err, messageOut).Error(), nil)
			return
		}
		h.logger.Info(logger.COSEC, logger.HandleReply, messageOut, nil)
	}()
	return h.Orchestrator.HandleReply(ctx, reply)
}
