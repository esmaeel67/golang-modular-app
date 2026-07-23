package handlers

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/cosec/internal/models"
	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/sec"
)

func RegisterReplyHandlers(subscriber am.ReplySubscribe, orchestrator sec.Orchestrator[*models.CreateOrderData]) error {
	replyMsgHandler := am.MessageHandlerFunc[am.IncomingReplyMessage](func(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
		return orchestrator.HandleReply(ctx, replyMsg)
	})

	return subscriber.Subscribe(orchestrator.ReplyTopic(), replyMsgHandler, am.GroupName("cosec-replies"))
}
