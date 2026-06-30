package queries

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
	"github.com/stackus/errors"
)

type GetOrderQuery struct {
	ID string
}

type GetOrderHandler struct {
	repo domain.OrderRepository
}

func NewGetOrderHandler(repo domain.OrderRepository) GetOrderHandler {
	return GetOrderHandler{
		repo: repo,
	}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrderQuery) (*domain.Order, error) {
	order, err := h.repo.Find(ctx, query.ID)

	return order, errors.Wrap(err, "get order query")
}
