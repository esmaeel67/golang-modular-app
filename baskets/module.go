package baskets

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/baskets/internal/application"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/handlers"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/logging"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	baskets := postgres.NewBasketRepository("baskets", mono.DB())

	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())

	if err != nil {
		return err
	}

	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	app := logging.LogApplicationAccess(
		application.New(baskets, stores, products, orders, domainDispatcher),
		mono.Logger(),
	)
	orderHandlers := logging.LogDomainEventHandlerAccess(
		application.NewOrderHandler(orders),
		mono.Logger(),
	)

	// Print the type of registrar
	fmt.Printf("🔍 Registrar type: %T\n", mono.RPC())
	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)

	return nil
}
