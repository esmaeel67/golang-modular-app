package internal

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/cosec/internal/models"
	"github.com/esmaeel67/golang-modular-app/customers/customerspb"
	"github.com/esmaeel67/golang-modular-app/depot/depotpb"
	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/esmaeel67/golang-modular-app/internal/sec"
)

const CreateOrderSagaName = "cosec.CreateOrder"
const CreateOrderReplyChannel = "mallbots.cosec.replies.CreateOrder"

type createOrderSaga struct {
	sec.Saga[*models.CreateOrderData]
}

func NewCreateOrderSaga() sec.Saga[*models.CreateOrderData] {
	saga := createOrderSaga{
		Saga: sec.NewSaga[*models.CreateOrderData](CreateOrderSagaName, CreateOrderReplyChannel),
	}
	// 0, -RejectOrder
	saga.AddStep().Compensation(saga.rejectOrder)

	// 1. AuthorizeCustomer
	saga.AddStep().Action(saga.authorizeCustomer)

	// 2. createShoppingList, -CancelShoppingList

	// 3. ConfirmPayment

	// 4. InitiateShopping

	// 5. ApproveOrder

	return saga
}

func (s createOrderSaga) rejectOrder(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(customerspb.AuthorizeCustomerCommand, customerspb.CommandChannel, &customerspb.AuthorizeCustomer{Id: data.CustomerID})
}

func (s createOrderSaga) authorizeCustomer(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(customerspb.AuthorizeCustomerCommand, customerspb.CommandChannel, &customerspb.AuthorizeCustomer{Id: data.CustomerID})
}
func (s createOrderSaga) createShoppingList(ctx context.Context, data *models.CreateOrderData) am.Command {
	items := make([]*depotpb.CreateShoppingList_Item, len(data.Items))
	for i, item := range data.Items {
		items[i] = &depotpb.CreateShoppingList_Item{
			ProductId: item.ProductID,
			StoreId:   item.StoreID,
			Quantity:  int32(item.Quantity),
		}
	}
	return am.NewCommand(depotpb.CreateShoppingListCommand, depotpb.CommandChannel, &depotpb.CreateShoppingList{
		OrderId: data.OrderID,
		Items:   items,
	})
}
