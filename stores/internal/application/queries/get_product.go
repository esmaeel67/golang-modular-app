package queries

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type GetProductQuery struct {
	ID string
}

type GetProductHandler struct {
	products domain.ProductRepository
}

func NewGetProductHandler(products domain.ProductRepository) GetProductHandler {
	return GetProductHandler{products: products}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProductQuery) (*domain.Product, error) {
	return h.products.FindProduct(ctx, query.ID)
}
