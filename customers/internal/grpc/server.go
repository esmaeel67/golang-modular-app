package grpc

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/customers/customerspb"
	pb "github.com/esmaeel67/golang-modular-app/customers/customerspb"
	"github.com/esmaeel67/golang-modular-app/customers/internal/application"
	"github.com/esmaeel67/golang-modular-app/customers/internal/domain"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	pb.UnimplementedCustomersServiceServer
}

var _ pb.CustomersServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	pb.RegisterCustomersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) RegisterCustomer(ctx context.Context, request *pb.RegisterCustomerRequest) (*pb.RegisterCustomerResponse, error) {
	id := uuid.New().String()
	err := s.app.RegisterCustomer(ctx, application.RegisterCustomer{
		ID:        id,
		Name:      request.GetName(),
		SmsNumber: request.GetSmsNumber(),
	})

	return &pb.RegisterCustomerResponse{Id: id}, err
}

func (s server) AuthorizeCustomer(ctx context.Context, request *pb.AuthorizeCustomerRequest) (*pb.AuthorizeCustomerResponse, error) {
	err := s.app.AuthorizeCustomer(ctx, application.AuthorizeCustomer{ID: request.GetId()})
	if err != nil {
		return nil, err
	}
	return &pb.AuthorizeCustomerResponse{}, nil
}

func (s server) GetCustomer(ctx context.Context, request *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customer, err := s.app.GetCustomer(ctx, application.GetCustomer{ID: request.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetCustomerResponse{
		Customer: s.customerFromDomain(customer),
	}, nil
}

func (s server) EnableCustomer(ctx context.Context, request *pb.EnableCustomerRequest) (*pb.EnableCustomerResponse, error) {
	err := s.app.EnableCustomer(ctx, application.EnableCustomer{ID: request.GetId()})
	return &pb.EnableCustomerResponse{}, err
}
func (s server) DisableCustomer(ctx context.Context, request *pb.DisableCustomerRequest) (*pb.DisableCustomerResponse, error) {
	err := s.app.DisableCustomer(ctx, application.DisableCustomer{ID: request.GetId()})
	return &pb.DisableCustomerResponse{}, err
}

func (s server) customerFromDomain(customer *domain.Customer) *customerspb.Customer {
	return &customerspb.Customer{
		Id:        customer.ID(),
		Name:      customer.Name,
		SmsNumber: customer.SmsNumber,
		Enabled:   customer.Enabled,
	}
}
