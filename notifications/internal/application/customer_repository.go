package application

import (
	"context"

	models "github.com/esmaeel67/golang-modular-app/notifications/internal/model"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
