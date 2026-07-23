package handlers

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/cosec/internal/models"
	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/sec"
	"github.com/esmaeel67/golang-modular-app/ordering/orderingpb"
)

type integrationEventHandlers[T ddd.Event] struct {
	orchestrator sec.Orchestrator[*models.CreateOrderData]
}

var _ ddd.EventHandler[ddd.Event] = (*integrationEventHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(saga sec.Orchestrator[*models.CreateOrderData]) ddd.EventHandler[ddd.Event] {
	return integrationEventHandlers[ddd.Event]{
		orchestrator: saga,
	}
}

func RegisterIntegrationEventHandlers(subscriber am.EventSubscriber, handlers ddd.EventHandler[ddd.Event]) (err error) {

	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return handlers.HandleEvent(ctx, eventMsg)
	})

	return subscriber.Subscribe(orderingpb.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderingpb.OrderCanceledEvent,
	}, am.GroupName("cosec-ordering"))

}

func (h integrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {

	switch event.EventName() {
	case orderingpb.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	}
	return nil
}

func (h integrationEventHandlers[T]) onOrderCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*orderingpb.OrderCreated)

	var total float64
	items := make([]models.Item, len(payload.GetItems()))
	for i, item := range payload.GetItems() {
		items[i] = models.Item{
			ProductID: item.GetProductId(),
			StoreID:   item.GetStoreId(),
			Price:     item.GetPrice(),
			Quantity:  int(item.GetQuantity()),
		}
		total += float64(item.GetQuantity() * int32(item.GetPrice()))
	}

	data := &models.CreateOrderData{
		OrderID:    payload.GetId(),
		CustomerID: payload.GetCustomerId(),
		PaymentID:  payload.GetPaymentId(),
		Items:      items,
		Total:      total,
	}

	// Start the CreateOrderSaga
	return h.orchestrator.Start(ctx, event.ID(), data)

}
