package handlers

import (
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscribe ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscribe.Subscribe(invoiceHandlers, domain.OrderReadiedEvent)
}
