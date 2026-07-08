package handlers

import (
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

func RegisterCatalogHandlers(catalogHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(catalogHandlers,
		domain.ProductAddedEvent,
		domain.ProductRebrandedEvent,
		domain.ProductPriceIncreasedEvent,
		domain.ProductPriceDecreasedEvent,
		domain.ProductRemovedEvent,
	)
	// domainSubscriber.Subscribe(domain.ProductRebrandedEvent, catalogHandlers)
	// domainSubscriber.Subscribe(domain.ProductPriceIncreasedEvent, catalogHandlers)
	// domainSubscriber.Subscribe(domain.ProductPriceDecreasedEvent, catalogHandlers)
	// domainSubscriber.Subscribe(domain.ProductRemovedEvent, catalogHandlers)

}
