package queries

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type GetCatalogQuery struct {
	StoreID string
}

type GetCatalogHandler struct {
	products domain.ProductRepository
}

func NewGetCatalogHandler(products domain.ProductRepository) GetCatalogHandler {
	return GetCatalogHandler{products: products}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalogQuery) ([]*domain.Product, error) {
	return h.products.GetCatalog(ctx, query.StoreID)
}
