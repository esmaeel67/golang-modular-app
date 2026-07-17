package grpc

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/customers/customerspb"
	"github.com/esmaeel67/golang-modular-app/search/internal/application"
	"github.com/esmaeel67/golang-modular-app/search/internal/models"
	"google.golang.org/grpc"
)

type CustomerRepository struct {
	client customerspb.CustomersServiceClient
}

var _ application.CustomerRepository = (*CustomerRepository)(nil)

func NewCustomerRepository(conn *grpc.ClientConn) CustomerRepository {
	return CustomerRepository{
		client: customerspb.NewCustomersServiceClient(conn),
	}
}

func (r CustomerRepository) Find(ctx context.Context, customerID string) (*models.Customer, error) {
	resp, err := r.client.GetCustomer(ctx, &customerspb.GetCustomerRequest{Id: customerID})
	if err != nil {
		return nil, err
	}

	return &models.Customer{
		ID:   resp.Customer.GetId(),
		Name: resp.Customer.GetName(),
	}, nil
}
