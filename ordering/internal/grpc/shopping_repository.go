package grpc

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/depotpb"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/domain"
	"google.golang.org/grpc"
)

type ShoppingRepository struct {
	client depotpb.DepotServiceClient
}

var _ domain.ShoppingRepository = (*ShoppingRepository)(nil)

func NewShoppingRepository(conn *grpc.ClientConn) ShoppingRepository {
	return ShoppingRepository{client: depotpb.NewDepotServiceClient(conn)}
}

func (r ShoppingRepository) Create(ctx context.Context, orderID string, orderItems []domain.Item) (string, error) {
	items := make([]*depotpb.OrderItem, 0, len(orderItems))

	for i, item := range orderItems {
		items[i] = r.itemFromDomain(item)
	}

	response, err := r.client.CreateShoppingList(ctx, &depotpb.CreateShoppingListRequest{
		OrderId: orderID,
		Items:   items,
	})
	if err != nil {
		return "", err
	}

	return response.GetId(), nil
}

func (r ShoppingRepository) Cancel(ctx context.Context, shoppingID string) error {
	_, err := r.client.CancelShoppingList(ctx, &depotpb.CancelShoppingListRequest{Id: shoppingID})
	return err
}

func (r ShoppingRepository) itemFromDomain(item domain.Item) *depotpb.OrderItem {
	return &depotpb.OrderItem{
		ProductId: item.ProductID,
		StoreId:   item.StoreID,
		Quantity:  int32(item.Quantity),
	}
}
