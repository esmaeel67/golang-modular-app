package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type RemoveProductCommand struct {
	ID string
}

type RemoveProductHandler struct {
	stores   domain.StoreRepository
	products domain.ProductRepository
}

func NewRemoveProductHandler(stores domain.StoreRepository, products domain.ProductRepository) RemoveProductHandler {
	return RemoveProductHandler{stores: stores, products: products}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProductCommand) error {
	return h.products.RemoveProduct(ctx, cmd.ID)
}
