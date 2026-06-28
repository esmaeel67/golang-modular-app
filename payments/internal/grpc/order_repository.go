package grpc

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/payments/internal/application"
	googlegrpc "google.golang.org/grpc"
)

type OrderRepository struct {
	// client orderingpb.OrderingServiceClient
}

var _ application.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(conn *googlegrpc.ClientConn) OrderRepository {
	return OrderRepository{
		// client: orderingpb.NewOrderingServiceClient(conn),
	}
}

func (r OrderRepository) Complete(ctx context.Context, invoiceID, orderID string) error {
	//TODO: fix client request
	// _, err := r.client.CompleteOrder(ctx, &orderingpb.CompleteOrderRequest{
	// 	Id:        orderID,
	// 	InvoiceId: invoiceID,
	// })
	// return err
	return nil
}
