package handlers

import (
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

func RegisterNotificationHandlers(notificationHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderCreated{}, notificationHandlers.OnOrderCreated)
	domainSubscriber.Subscribe(domain.OrderReadied{}, notificationHandlers.OnOrderReadied)
	domainSubscriber.Subscribe(domain.OrderCanceled{}, notificationHandlers.OnOrderCanceled)
}
