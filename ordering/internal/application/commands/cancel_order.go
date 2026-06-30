package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

type CancelOrderCommand struct {
	ID string
}

type CancelOrderHandler struct {
	orders          domain.OrderRepository
	shopping        domain.ShoppingRepository
	domainPublisher ddd.EventPublisher
}

func NewCancelOrderHandler(orders domain.OrderRepository, shopping domain.ShoppingRepository, domainPublisher ddd.EventPublisher) CancelOrderHandler {
	return CancelOrderHandler{
		orders:          orders,
		shopping:        shopping,
		domainPublisher: domainPublisher,
	}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrderCommand) error {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Cancel(); err != nil {
		return err
	}

	if err = h.shopping.Cancel(ctx, order.ShoppingID); err != nil {
		return err
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return err
	}

	// publish domain event
	if err = h.domainPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return err
	}

	return nil
}
