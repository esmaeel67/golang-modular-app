package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

type ReadyOrderCommand struct {
	ID string
}

type ReadyOrderHandler struct {
	orders domain.OrderRepository
}

func NewReadyOrderHandler(orders domain.OrderRepository) ReadyOrderHandler {
	return ReadyOrderHandler{
		orders: orders,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrderCommand) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Ready(); err != nil {
		return nil
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	return nil
}
