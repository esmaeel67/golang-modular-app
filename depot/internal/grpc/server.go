package grpc

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/depotpb"
	"github.com/esmaeel67/golang-modular-app/depot/internal/application"
	"github.com/esmaeel67/golang-modular-app/depot/internal/application/commands"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	depotpb.UnimplementedDepotServiceServer
}

var _ depotpb.DepotServiceServer = (*server)(nil)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	depotpb.RegisterDepotServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateShoppingList(ctx context.Context, request *depotpb.CreateShoppingListRequest) (*depotpb.CreateShoppingListResponse, error) {
	id := uuid.New().String()

	items := make([]commands.OrderItem, 0, len(request.GetItems()))
	for _, item := range request.GetItems() {
		items = append(items, s.itemToDomain(item))
	}

	err := s.app.CreateShoppingList(ctx, commands.CreateShoppingListCommand{
		ID:      id,
		OrderID: request.GetOrderId(),
		Items:   items,
	})

	return &depotpb.CreateShoppingListResponse{Id: id}, err
}

func (s server) CancelShoppingList(ctx context.Context, request *depotpb.CancelShoppingListRequest) (*depotpb.CancelShoppingListResponse, error) {
	err := s.app.CancelShoppingList(ctx, commands.CancelShoppingListCommand{
		ID: request.GetId(),
	})

	return &depotpb.CancelShoppingListResponse{}, err
}

func (s server) AssignShoppingList(ctx context.Context, request *depotpb.AssignShoppingListRequest) (*depotpb.AssignShoppingListResponse, error) {
	err := s.app.AssignShoppingList(ctx, commands.AssignShoppingListCommand{
		ID:    request.GetId(),
		BotID: request.GetBotId(),
	})
	return &depotpb.AssignShoppingListResponse{}, err
}

func (s server) CompleteShoppingList(ctx context.Context, request *depotpb.CompleteShoppingListRequest) (*depotpb.CompleteShoppingListResponse, error) {
	err := s.app.CompleteShoppingList(ctx, commands.CompleteShoppingListCommand{ID: request.GetId()})
	return &depotpb.CompleteShoppingListResponse{}, err
}

func (s server) itemToDomain(item *depotpb.OrderItem) commands.OrderItem {
	return commands.OrderItem{
		StoreID:   item.GetStoreId(),
		ProductID: item.GetProductId(),
		Quantity:  int(item.GetQuantity()),
	}
}
