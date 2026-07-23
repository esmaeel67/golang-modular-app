package paymentspb

import (
	"github.com/esmaeel67/golang-modular-app/internal/registry"
	"github.com/esmaeel67/golang-modular-app/internal/registry/serdes"
)

const (
	InvoiceAggregateChannel = "mallbots.payments.events.Invoice"

	InvoicePaidEvent = "paymentsapi.InvoicePaid"

	CommandChannel = "mallbots.payments.commands"

	ConfirmPaymentCommand = "paymentsapi.ConfirmPayment"
)

func Registrations(reg registry.Registry) (err error) {
	serde := serdes.NewProtoSerde(reg)

	// Invoice events
	if err = serde.Register(&InvoicePaid{}); err != nil {
		return err
	}

	// commands
	if err = serde.Register(&ConfirmPayment{}); err != nil {
		return err
	}
	return
}

func (*InvoicePaid) Key() string {
	return InvoicePaidEvent
}

func (*ConfirmPayment) Key() string {
	return ConfirmPaymentCommand
}
