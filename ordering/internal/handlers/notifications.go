package handlers

import (
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

func RegisterNotificationHandlers(notificationHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(notificationHandlers, domain.OrderCreatedEvent, domain.OrderReadiedEvent, domain.OrderCanceledEvent)
	// domainSubscriber.Subscribe(notificationHandlers)
	// domainSubscriber.Subscribe(domain.OrderCanceledEvent, notificationHandlers)
}
