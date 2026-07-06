package application

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/application/commands"
	"github.com/esmaeel67/golang-modular-app/stores/internal/application/queries"
	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateStore(ctx context.Context, cmd commands.CreateStore) error
		EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) error
		DisableParticipation(ctx context.Context, cmd commands.DisableParticipation) error
		AddProduct(ctx context.Context, cmd commands.AddProduct) error
		RemoveProduct(ctx context.Context, cmd commands.RemoveProductCommand) error
	}
	Queries interface {
		GetStore(ctx context.Context, query queries.GetStore) (*domain.MallStore, error)
		GetStores(ctx context.Context, query queries.GetStores) ([]*domain.MallStore, error)
		GetParticipatingStores(ctx context.Context, query queries.GetParticipatingStores) ([]*domain.MallStore, error)
		GetCatalog(ctx context.Context, query queries.GetCatalogQuery) ([]*domain.CatalogProduct, error)
		GetProduct(ctx context.Context, query queries.GetProductQuery) (*domain.CatalogProduct, error)
	}
	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateStoreHandler
		commands.EnableParticipationHandler
		commands.DisableParticipationHandler
		commands.AddProductHandler
		commands.RemoveProductHandler
	}
	appQueries struct {
		queries.GetStoreHandler
		queries.GetStoresHandler
		queries.GetParticipatingStoriesHandler
		queries.GetCatalogHandler
		queries.GetProductHandler
	}
)

var _ App = (*Application)(nil)

func New(stores domain.StoreRepository, products domain.ProductRepository,
	catalog domain.CatalogRepository, mall domain.MallRepository,
) *Application {
	return &Application{
		appCommands: appCommands{
			CreateStoreHandler:          commands.NewCreateStoreHandler(stores),
			EnableParticipationHandler:  commands.NewEnableParticipationHandler(stores),
			DisableParticipationHandler: commands.NewDisableParticipationHandler(stores),
			AddProductHandler:           commands.NewAddProductHandler(products),
			RemoveProductHandler:        commands.NewRemoveProductHandler(products),
		},
		appQueries: appQueries{
			GetStoreHandler:                queries.NewGetStoreHandler(mall),
			GetStoresHandler:               queries.NewGetStoresHandler(mall),
			GetParticipatingStoriesHandler: queries.NewGetParticipatingStoresHandler(mall),
			GetCatalogHandler:              queries.NewGetCatalogHandler(catalog),
			GetProductHandler:              queries.NewGetProductHandler(catalog),
		},
	}
}
