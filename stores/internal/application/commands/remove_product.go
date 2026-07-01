package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type RemoveProductCommand struct {
	ID string
}

type RemoveProductHandler struct {
	stores          domain.StoreRepository
	products        domain.ProductRepository
	domainPublisher ddd.EventPublisher
}

func NewRemoveProductHandler(stores domain.StoreRepository, products domain.ProductRepository, domainPublisher ddd.EventPublisher) RemoveProductHandler {
	return RemoveProductHandler{stores: stores, products: products, domainPublisher: domainPublisher}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProductCommand) error {
	product, err := h.products.FindProduct(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = product.Remove(); err != nil {
		return err
	}

	if err = h.products.RemoveProduct(ctx, cmd.ID); err != nil {
		return err
	}

	// publish domain events
	if err = h.domainPublisher.Publish(ctx, product.GetEvents()...); err != nil {
		return err
	}

	return nil
}
