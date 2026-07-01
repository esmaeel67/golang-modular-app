package stores

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/stores/internal/application"
	"github.com/esmaeel67/golang-modular-app/stores/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/stores/internal/logging"
	"github.com/esmaeel67/golang-modular-app/stores/internal/postgres"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	domainDispatcher := ddd.NewEventDispatcher()

	stores := postgres.NewStoreRepository("stores", mono.DB())
	participatingStores := postgres.NewParticipatingStoryRepository("stores", mono.DB())
	products := postgres.NewProductRepository("products", mono.DB())

	var app application.App
	app = application.New(stores, participatingStores, products, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterService(ctx, app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
