package grpc

import (
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application/commands"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
	"github.com/esmaeel67/golang-modular-app/ordering/orderingpb"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	orderingpb.UnimplementedOrderingServiceServer
}

var _ orderingpb.OrderingServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {

	orderingpb.RegisterOrderingServiceServer(registrar, server{app: app})

	return nil
}

func (s server) CreateOrder(ctx context.Context, request *orderingpb.CreateOrderRequest) (*orderingpb.CreateOrderResponse, error) {
	id := uuid.New().String()

	items := make([]domain.Item, 0, len(request.Items))
	for _, item := range request.Items {
		items = append(items, s.itemToDomain(item))
	}

	err := s.app.CreateOrder(ctx, commands.CreateOrderCommand{
		ID:         id,
		CustomerID: request.GetCustomerId(),
		PaymentID:  request.GetPaymentId(),
		Items:      items,
	})

	return &orderingpb.CreateOrderResponse{Id: id}, err

}

func (s server) itemToDomain(item *orderingpb.Item) domain.Item {
	return domain.Item{
		ProductID:   item.GetProductId(),
		StoreID:     item.GetStoreId(),
		StoreName:   item.GetStoreName(),
		ProductName: item.GetProductName(),
		Price:       item.GetPrice(),
		Quantity:    int(item.GetQuantity()),
	}
}

func (s server) itemFromDomain(item *domain.Item) *orderingpb.Item {
	return &orderingpb.Item{
		StoreId:     item.StoreID,
		ProductId:   item.ProductID,
		StoreName:   item.StoreName,
		ProductName: item.ProductName,
		Price:       item.Price,
		Quantity:    int32(item.Quantity),
	}
}
