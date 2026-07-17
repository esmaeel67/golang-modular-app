package handlers

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/customers/customerspb"
	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
)

func RegisterCustomerHandlers(customerHandlers ddd.EventHandler[ddd.Event], stream am.EventSubscriber) error {
	evtMsgHandler := am.MessageHandlerFunc[am.EventMessage](func(ctx context.Context, eventMsg am.EventMessage) error {
		return customerHandlers.HandleEvent(ctx, eventMsg)
	})

	return stream.Subscribe(customerspb.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		customerspb.CustomerRegisteredEvent,
	}, am.GroupName("search-customers"))
}
