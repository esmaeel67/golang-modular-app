package application

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/baskets/internal/domain"
	"github.com/stackus/errors"
)

type StartBasket struct {
	ID         string
	CustomerID string
}

type CancelBasket struct {
	ID string
}

type CheckoutBasket struct {
	ID        string
	PaymentID string
}

type AddItem struct {
	ID        string
	ProductID string
	Quantity  int
}

type RemoveItem struct {
	ID        string
	ProductID string
	Quantity  int
}
type GetBasket struct {
	ID string
}

type App interface {
	StartBasket(ctx context.Context, start StartBasket) error
	CancelBasket(ctx context.Context, cancel CancelBasket) error
	CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error
	AddItem(ctx context.Context, add AddItem) error
	RemoveItem(ctx context.Context, remove RemoveItem) error
	GetBasket(ctx context.Context, get GetBasket) (*domain.Basket, error)
}

type Application struct {
	baskets domain.BasketRepository
}

var _ App = (*Application)(nil)

func New(baskets domain.BasketRepository) *Application {
	if baskets == nil {
		panic("baskets repository cannot be nil")
	}

	return &Application{
		baskets: baskets,
	}
}

func (a *Application) StartBasket(ctx context.Context, start StartBasket) error {
	basket, err := domain.StartBasket(start.ID, start.CustomerID)

	if err != nil {
		return err
	}
	fmt.Println(basket)
	return a.baskets.Save(ctx, basket)
}

func (a *Application) CancelBasket(ctx context.Context, cancel CancelBasket) error {
	basket, err := a.baskets.Find(ctx, cancel.ID)

	if err != nil {
		return err
	}

	err = basket.Cancel()
	if err != nil {
		return err
	}

	return a.baskets.Update(ctx, basket)
}
func (a *Application) CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error {
	basket, err := a.baskets.Find(ctx, checkout.ID)
	if err != nil {
		return err
	}
	err = basket.Checkout(checkout.PaymentID)
	if err != nil {
		return errors.Wrap(err, "basket checkout")
	}

	// submit the basket to the order module
	// _,err = a.order.save(ctx,basket)
	// if err != nil {
	// 	return errors.Wrap(err,"baskets checkout")
	// }

	return errors.Wrap(a.baskets.Update(ctx, basket), "basket checkout")

}

func (a *Application) AddItem(ctx context.Context, add AddItem) error {
	return nil
}

func (a *Application) RemoveItem(ctx context.Context, remove RemoveItem) error {
	return nil
}
func (a *Application) GetBasket(ctx context.Context, get GetBasket) (*domain.Basket, error) {
	return a.baskets.Find(ctx, get.ID)
}
