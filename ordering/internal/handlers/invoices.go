package handlers

import (
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscribe ddd.EventSubscriber) {
	domainSubscribe.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
