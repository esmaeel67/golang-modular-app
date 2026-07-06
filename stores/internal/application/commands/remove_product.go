package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type RemoveProductCommand struct {
	ID string
}

type RemoveProductHandler struct {
	products domain.ProductRepository
}

func NewRemoveProductHandler(products domain.ProductRepository) RemoveProductHandler {
	return RemoveProductHandler{products: products}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProductCommand) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.Remove(); err != nil {
		return err
	}

	return h.products.Save(ctx, product)
}
