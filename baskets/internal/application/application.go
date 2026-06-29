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
	baskets  domain.BasketRepository
	stores   domain.StoreRepository
	products domain.ProductRepository
	orders   domain.OrderRepository
}

var _ App = (*Application)(nil)

func New(baskets domain.BasketRepository, stores domain.StoreRepository, products domain.ProductRepository, orders domain.OrderRepository) *Application {
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
	_, err = a.orders.Save(ctx, basket)
	if err != nil {
		return errors.Wrap(err, "baskets checkout")
	}

	return errors.Wrap(a.baskets.Update(ctx, basket), "basket checkout")

}

func (a *Application) AddItem(ctx context.Context, add AddItem) error {
	basket, err := a.baskets.Find(ctx, add.ID)
	if err != nil {
		return err
	}

	product, err := a.products.Find(ctx, add.ProductID)
	if err != nil {
		return err
	}

	store, err := a.stores.Find(ctx, product.StoreID)
	if err != nil {
		return err
	}

	err = basket.AddItem(store, product, add.Quantity)
	if err != nil {
		return err
	}
	return a.baskets.Update(ctx, basket)
}

func (a *Application) RemoveItem(ctx context.Context, remove RemoveItem) error {
	product, err := a.products.Find(ctx, remove.ProductID)
	if err != nil {
		return err
	}

	basket, err := a.baskets.Find(ctx, remove.ID)
	if err != nil {
		return err
	}

	err = basket.RemoveItem(product, remove.Quantity)
	if err != nil {
		return nil
	}

	return a.baskets.Update(ctx, basket)
}
func (a *Application) GetBasket(ctx context.Context, get GetBasket) (*domain.Basket, error) {
	return a.baskets.Find(ctx, get.ID)
}
