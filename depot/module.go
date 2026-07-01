package depot

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/depot/internal/application"
	"github.com/esmaeel67/golang-modular-app/depot/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/depot/internal/handlers"
	"github.com/esmaeel67/golang-modular-app/depot/internal/logging"
	"github.com/esmaeel67/golang-modular-app/depot/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	// setup Driven adapter
	domainDispatcher := ddd.NewEventDispatcher()
	shoppingLists := postgres.NewShoppingListRepository("shopping_lists", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(
		application.New(shoppingLists, stores, products, domainDispatcher),
		mono.Logger())
	orderHandlers := logging.LogDomainEventHandlerAccess(
		application.NewOrderHandlers(orders),
		mono.Logger(),
	)

	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}

	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)

	return nil
}
