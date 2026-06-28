package grpc

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/notifications/internal/application"
	pb "github.com/esmaeel67/golang-modular-app/notifications/notificationspb"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	pb.UnimplementedNotificationsServiceServer
}

var _ pb.NotificationsServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	pb.RegisterNotificationsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) NotifyOrderCreated(ctx context.Context, request *pb.NotifyOrderCreatedRequest) (*pb.NotifyOrderCreatedResponse, error) {
	err := s.app.NotifyOrderCreated(ctx, application.OrderCreated{
		OrderID:    request.GetOrderId(),
		CustomerID: request.GetCustomerId(),
	})
	return &pb.NotifyOrderCreatedResponse{}, err
}

func (s server) NotifyOrderCanceled(ctx context.Context, request *pb.NotifyOrderCanceledRequest) (*pb.NotifyOrderCanceledResponse, error) {
	err := s.app.NotifyOrderCanceled(ctx, application.OrderCanceled{
		OrderID:    request.GetOrderId(),
		CustomerID: request.GetCustomerId(),
	})
	return &pb.NotifyOrderCanceledResponse{}, err
}

func (s server) NotifyOrderReady(ctx context.Context, request *pb.NotifyOrderReadyRequest) (*pb.NotifyOrderReadyResponse, error) {
	err := s.app.NotifyOrderReady(ctx, application.OrderReady{
		OrderID:    request.GetOrderId(),
		CustomerID: request.GetCustomerId(),
	})
	return &pb.NotifyOrderReadyResponse{}, err
}
