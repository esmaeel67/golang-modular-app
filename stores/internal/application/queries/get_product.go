package queries

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type GetProductQuery struct {
	ID string
}

type GetProductHandler struct {
	catalog domain.CatalogRepository
}

func NewGetProductHandler(catalog domain.CatalogRepository) GetProductHandler {
	return GetProductHandler{catalog: catalog}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProductQuery) (*domain.CatalogProduct, error) {
	return h.catalog.Find(ctx, query.ID)
}
