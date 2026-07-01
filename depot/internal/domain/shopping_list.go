package domain

import (
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/stackus/errors"
)

var (
	ErrShoppingCannotBeCancelled = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be cancelled")
	ErrShoppingCannotBeAssigned  = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot bet assigned")
	ErrShoppingCannotBeCompleted = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot bet completed")
)

type ShoppingList struct {
	ddd.AggregateBase
	OrderID       string
	Stops         Stops
	AssignedBotID string
	Status        ShoppingListStatus
}

func CreateShopping(id, orderID string) *ShoppingList {
	shoppingList := &ShoppingList{
		AggregateBase: ddd.AggregateBase{
			ID: id,
		},
		OrderID: orderID,
		Status:  ShoppingListIsAvailable,
		Stops:   make(Stops),
	}

	shoppingList.AddEvent(&ShoppingListCreated{
		ShoppingList: shoppingList,
	})

	return shoppingList
}

func (sl *ShoppingList) AddItem(store *Store, product *Product, quantity int) error {
	if _, exists := sl.Stops[store.ID]; !exists {
		sl.Stops[store.ID] = &Stop{
			StoreName:     store.Name,
			StoreLocation: store.Location,
			Items:         make(Items),
		}
	}
	return sl.Stops[store.ID].AddItem(product, quantity)
}

func (sl ShoppingList) isCancelable() bool {
	switch sl.Status {
	case ShoppingListIsAvailable, ShoppingListIsAssigned, ShoppingListIsActive:
		return true
	default:
		return false
	}
}

func (sl *ShoppingList) Cancel() error {
	if !sl.isCancelable() {
		return ErrShoppingCannotBeCancelled
	}

	sl.Status = ShoppingListIsCancelled

	sl.AddEvent(&ShoppingListCanceled{
		ShoppingList: sl,
	})

	return nil
}

func (sl ShoppingList) isAssignable() bool {
	return sl.Status == ShoppingListIsAvailable
}

func (sl *ShoppingList) Assign(id string) error {
	// validate status

	if !sl.isAssignable() {
		return ErrShoppingCannotBeAssigned
	}

	sl.AssignedBotID = id
	sl.Status = ShoppingListIsAssigned

	sl.AddEvent(&ShoppingListAssigned{
		ShoppingList: sl,
		BotID:        id,
	})

	return nil
}

func (sl ShoppingList) isCompletable() bool {
	return sl.Status == ShoppingListIsAssigned
}

func (sl *ShoppingList) Complete() error {
	if !sl.isCompletable() {
		return ErrShoppingCannotBeCompleted
	}

	sl.Status = ShoppingListIsCompleted

	sl.AddEvent(&ShoppingListCompleted{
		ShoppingList: sl,
	})

	return nil
}
