package queries

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type GetCatalogQuery struct {
	StoreID string
}

type GetCatalogHandler struct {
	catalog domain.CatalogRepository
}

func NewGetCatalogHandler(catalog domain.CatalogRepository) GetCatalogHandler {
	return GetCatalogHandler{catalog: catalog}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalogQuery) ([]*domain.CatalogProduct, error) {
	return h.catalog.GetCatalog(ctx, query.StoreID)
}
