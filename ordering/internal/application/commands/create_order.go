package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
)

type CreateOrderCommand struct {
	ID         string
	CustomerID string
	PaymentID  string
	Items      []*domain.Item
}

type CreateOrderHandler struct {
	orders        domain.OrderRepository
	customers     domain.CustomerRepository
	payments      domain.PaymentRepository
	shopping      domain.ShoppingRepository
	notifications domain.NotificationRepository
}

func NewCreateOrderHandler(orders domain.OrderRepository, customers domain.CustomerRepository,
	payments domain.PaymentRepository, shopping domain.ShoppingRepository, notifications domain.NotificationRepository,
) CreateOrderHandler {
	return CreateOrderHandler{
		orders:        orders,
		customers:     customers,
		payments:      payments,
		shopping:      shopping,
		notifications: notifications,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrderCommand) error {

	return nil
}
