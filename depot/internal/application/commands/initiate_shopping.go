package commands

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/internal/domain"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
)

type InitiateShoppingCommand struct {
	ID string
}

type InitiateShoppingHandler struct {
	shoppingLists   domain.ShoppingListRepository
	domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
}

func NewInitiateShoppingHandler(lists domain.ShoppingListRepository, publisher ddd.EventPublisher[ddd.AggregateEvent]) InitiateShoppingHandler {
	return InitiateShoppingHandler{
		shoppingLists:   lists,
		domainPublisher: publisher,
	}
}

func (h InitiateShoppingHandler) InitiateShopping(ctx context.Context, cmd InitiateShoppingCommand) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = list.Initiate()
	if err != nil {
		return err
	}

	if err = h.shoppingLists.Update(ctx, list); err != nil {
		return err
	}

	// publish domain event
	if err = h.domainPublisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
