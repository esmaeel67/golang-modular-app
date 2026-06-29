package depot

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/internal/application"
	"github.com/esmaeel67/golang-modular-app/depot/internal/logging"
	"github.com/esmaeel67/golang-modular-app/depot/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/depot/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	// setup Driven adapter
	shoppingLists := postgres.NewShoppingListRepository("shopping_lists", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	var app application.App
	app = application.New(shoppingLists, stores, products, orders)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
