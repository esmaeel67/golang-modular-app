package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/internal/domain"
)

type CreateShoppingListCommand struct {
	ID      string
	OrderID string
	Items   []OrderItem
}

type CreateShoppingListHandler struct {
	shoppingLists domain.ShoppingListRepository
	stores        domain.StoreRepository
	products      domain.ProductRepository
}

func NewCreateShoppingListHandler(shoppingLists domain.ShoppingListRepository, stores domain.StoreRepository, products domain.ProductRepository) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingLists: shoppingLists,
		stores:        stores,
		products:      products,
	}

}
func (h CreateShoppingListHandler) CreateShoppingList(ctx context.Context, cmd CreateShoppingListCommand) error {

	return nil
}
