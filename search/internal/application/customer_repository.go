package application

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/search/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
