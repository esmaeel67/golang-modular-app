package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/internal/domain"
)

type AssignShoppingListCommand struct {
	ID    string
	BotID string
}

type AssignShoppingListHandler struct {
	shoppingLists domain.ShoppingListRepository
}

func NewAssignShoppingListHandler(shoppingList domain.ShoppingListRepository) AssignShoppingListHandler {
	return AssignShoppingListHandler{
		shoppingLists: shoppingList,
	}
}

func (h AssignShoppingListHandler) AssignShoppingList(ctx context.Context, cmd AssignShoppingListCommand) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Assign(cmd.BotID)
	if err != nil {
		return err
	}

	return h.shoppingLists.Update(ctx, list)
}
