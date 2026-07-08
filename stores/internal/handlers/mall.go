package handlers

import (
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

func RegisterMallHandlers(mallHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {

	domainSubscriber.Subscribe(mallHandlers,
		domain.StoreCreatedEvent,
		domain.StoreParticipationEnabledEvent,
		domain.StoreParticipationDisabledEvent,
		domain.StoreRebrandedEvent,
	)
	// domainSubscriber.Subscribe(domain.StoreParticipationEnabledEvent, mallHandlers)
	// domainSubscriber.Subscribe(domain.StoreParticipationDisabledEvent, mallHandlers)
	// domainSubscriber.Subscribe(domain.StoreRebrandedEvent, mallHandlers)

}
