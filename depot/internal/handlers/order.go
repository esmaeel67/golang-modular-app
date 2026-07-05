package handlers

import (
	"github.com/esmaeel67/golang-modular-app/depot/internal/domain"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.ShoppingListCompletedEvent, orderHandlers)
}
