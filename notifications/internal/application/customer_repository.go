package application

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
