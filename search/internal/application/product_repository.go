package application

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/search/internal/models"
)

type ProductRepository interface {
	Find(ctx context.Context, productID string) (*models.Product, error)
}
