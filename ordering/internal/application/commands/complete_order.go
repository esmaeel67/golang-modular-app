package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

type CompleteOrderCommand struct {
	ID        string
	InvoiceID string
}

type CompleteOrderHandler struct {
	orders domain.OrderRepository
}

func NewCompleteOrderHandler(orders domain.OrderRepository) CompleteOrderHandler {
	return CompleteOrderHandler{
		orders: orders,
	}
}

func (h CompleteOrderHandler) CompleteOrder(ctx context.Context, cmd CompleteOrderCommand) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = order.Complete(cmd.InvoiceID)
	if err != nil {
		return err
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	return nil
}
