package handlers

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/customers/customerspb"
	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/ordering/orderingpb"
	"github.com/esmaeel67/golang-modular-app/search/internal/application"
	"github.com/esmaeel67/golang-modular-app/stores/storespb"
)

type integrationHandlers[T ddd.Event] struct {
	orders    application.OrderRepository
	customers application.CustomerCacheRepository
	products  application.ProductCacheRepository
	stores    application.StoreCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(orders application.OrderRepository, customers application.CustomerCacheRepository,
	stores application.StoreCacheRepository, products application.ProductCacheRepository) ddd.EventHandler[ddd.Event] {
	return integrationHandlers[ddd.Event]{
		orders:    orders,
		customers: customers,
		products:  products,
		stores:    stores,
	}
}
func RegisterIntegrationEventHandlers(subscriber am.EventSubscriber, handlers ddd.EventHandler[ddd.Event]) (err error) {

	evtMsgHandler := am.MessageHandlerFunc[am.IncomingEventMessage](func(ctx context.Context, eventMsg am.IncomingEventMessage) error {
		return handlers.HandleEvent(ctx, eventMsg)
	})

	if err = subscriber.Subscribe(customerspb.CustomerAggregateChannel, evtMsgHandler, am.MessageFilter{
		customerspb.CustomerRegisteredEvent,
	}, am.GroupName("search-customers")); err != nil {
		return
	}

	if err = subscriber.Subscribe(orderingpb.OrderAggregateChannel, evtMsgHandler, am.MessageFilter{
		orderingpb.OrderCreatedEvent,
		orderingpb.OrderReadiedEvent,
		orderingpb.OrderCanceledEvent,
		orderingpb.OrderCompletedEvent,
	}, am.GroupName("notification-orders")); err != nil {
		return
	}

	if err = subscriber.Subscribe(storespb.ProductAggregateChannel, evtMsgHandler, am.MessageFilter{
		storespb.ProductAddedEvent,
		storespb.ProductRebrandedEvent,
		storespb.ProductRemovedEvent,
	}, am.GroupName("search-products")); err != nil {
		return
	}

	if err = subscriber.Subscribe(storespb.StoreAggregateChannel, evtMsgHandler, am.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreRebrandedEvent,
	}); err != nil {
		return
	}

	return
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) error {

	return nil
}
