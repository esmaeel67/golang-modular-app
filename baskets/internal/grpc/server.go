package grpc

import (
	"context"
	"fmt"

	pb "github.com/esmaeel67/golang-modular-app/baskets/basketspb"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/application"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/domain"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type BasketServer struct {
	pb.UnimplementedBasketServiceServer
	app application.App
}

var _ pb.BasketServiceServer = (*BasketServer)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	// Check if registrar is nil
	if registrar == nil {
		return fmt.Errorf("registrar is nil")
	}

	fmt.Printf("🔍 Registrar app: %T\n", app)

	// Print the type of registrar
	fmt.Printf("🔍 Registrar type: %T\n", registrar)
	pb.RegisterBasketServiceServer(registrar, &BasketServer{app: app})
	return nil
}

func (s BasketServer) StartBasket(ctx context.Context, request *pb.StartBasketRequest) (*pb.StartBasketResponse, error) {
	basketID := uuid.New().String()

	fmt.Println("requestrequestrequest : ", request)

	fmt.Printf("🔍 StartBasket app: %T\n", s.app)
	err := s.app.StartBasket(ctx, application.StartBasket{
		ID:         basketID,
		CustomerID: request.GetCustomerId(),
	})

	return &pb.StartBasketResponse{Id: basketID}, err
}

func (s BasketServer) CancelBasket(ctx context.Context, request *pb.CancelBasketRequest) (*pb.CancelBasketResponse, error) {
	err := s.app.CancelBasket(ctx, application.CancelBasket{
		ID: request.GetId(),
	})

	return &pb.CancelBasketResponse{}, err
}

func (s BasketServer) CheckoutBasket(ctx context.Context, request *pb.CheckoutBasketRequest) (*pb.CheckoutBasketResponse, error) {
	err := s.app.CheckoutBasket(ctx, application.CheckoutBasket{
		ID:        request.GetId(),
		PaymentID: request.GetPaymentId(),
	})
	return &pb.CheckoutBasketResponse{}, err
}

func (s BasketServer) AddItem(ctx context.Context, request *pb.AddItemRequest) (*pb.AddItemResponse, error) {

	err := s.app.AddItem(ctx, application.AddItem{
		ID:        request.GetId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})

	return &pb.AddItemResponse{}, err
}

func (s BasketServer) RemoveItem(ctx context.Context, request *pb.RemoveItemRequest) (*pb.RemoveItemResponse, error) {

	err := s.app.RemoveItem(ctx, application.RemoveItem{
		ID:        request.GetId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})

	return &pb.RemoveItemResponse{}, err
}
func (s BasketServer) GetBasket(ctx context.Context, request *pb.GetBasketRequest) (*pb.GetBasketResponse, error) {

	basket, err := s.app.GetBasket(ctx, application.GetBasket{
		ID: request.GetId(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.GetBasketResponse{
		Basket: s.basketFromDomain(basket),
	}, nil
}
func (s BasketServer) basketFromDomain(basket *domain.Basket) *pb.Basket {
	protoBasket := &pb.Basket{
		Id: basket.ID(),
	}
	protoBasket.Items = make([]*pb.Item, 0, len(basket.Items))
	for _, item := range basket.Items {
		protoBasket.Items = append(protoBasket.Items, &pb.Item{
			StoreId:      item.StoreID,
			StoreName:    item.StoreName,
			ProductId:    item.ProductID,
			ProductName:  item.ProductName,
			ProductPrice: item.ProductPrice,
			Quantity:     int32(item.Quantity),
		})
	}
	return protoBasket
}
