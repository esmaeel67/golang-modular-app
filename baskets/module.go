package baskets

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/baskets/internal/application"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/logging"
	"github.com/esmaeel67/golang-modular-app/baskets/internal/postgres"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {

	baskets := postgres.NewBasketRepository("baskets", mono.DB())

	// conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())

	// if err != nil {
	// 	return err
	// }
	//TODO: implemented stores and product and orders repository
	// sores := grpc.NewStoreRepository(con)

	// setup application
	var app application.App
	app = application.New(baskets)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// Print the type of registrar
	fmt.Printf("🔍 Registrar type: %T\n", mono.RPC())
	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
