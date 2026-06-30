package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

type ReadyOrderCommand struct {
	ID string
}

type ReadyOrderHandler struct {
	orders          domain.OrderRepository
	domainPublisher ddd.EventPublisher
}

func NewReadyOrderHandler(orders domain.OrderRepository, domainPublisher ddd.EventPublisher) ReadyOrderHandler {
	return ReadyOrderHandler{
		orders:          orders,
		domainPublisher: domainPublisher,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrderCommand) error {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Ready(); err != nil {
		return nil
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return err
	}

	// publish domain events
	if err = h.domainPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return err
	}

	return nil
}
