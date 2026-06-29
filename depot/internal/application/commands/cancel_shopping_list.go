package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/internal/domain"
)

type CancelShoppingListCommand struct {
	ID string
}

type CancelShoppingListHandler struct {
	shoppingLists domain.ShoppingListRepository
}

func NewCancelShoppingListHandler(shoppingLists domain.ShoppingListRepository) CancelShoppingListHandler {
	return CancelShoppingListHandler{
		shoppingLists: shoppingLists,
	}
}

func (h CancelShoppingListHandler) CancelShoppingList(ctx context.Context, cmd CancelShoppingListCommand) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Cancel()
	if err != nil {
		return err
	}

	return h.shoppingLists.Update(ctx, list)
}
