package grpc

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/application"
	"github.com/esmaeel67/golang-modular-app/stores/internal/application/commands"
	"github.com/esmaeel67/golang-modular-app/stores/internal/application/queries"
	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
	pb "github.com/esmaeel67/golang-modular-app/stores/storespb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	pb.UnimplementedStoresServiceServer
}

var _ pb.StoresServiceServer = (*server)(nil)

func RegisterService(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	pb.RegisterStoresServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateStore(ctx context.Context, request *pb.CreateStoreRequest) (*pb.CreateStoreResponse, error) {

	storeID := uuid.New().String()

	err := s.app.CreateStore(ctx, commands.CreateStore{
		ID:       storeID,
		Name:     request.GetName(),
		Location: request.GetLocation(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateStoreResponse{
		Id: storeID,
	}, nil
}

func (s server) GetStore(ctx context.Context, request *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	stores, err := s.app.GetStore(ctx, queries.GetStore{ID: request.GetId()})
	if err != nil {
		return nil, err
	}
	return &pb.GetStoreResponse{Store: s.storeFromDomain(stores)}, nil
}

func (s server) GetStores(ctx context.Context, request *pb.GetStoresRequest) (*pb.GetStoresResponse, error) {
	stores, err := s.app.GetStores(ctx, queries.GetStores{})
	if err != nil {
		return nil, err
	}
	protoStores := []*pb.Store{}

	for _, store := range stores {
		protoStores = append(protoStores, s.storeFromDomain(store))
	}
	return &pb.GetStoresResponse{
		Stores: protoStores,
	}, nil
}

func (s server) EnableParticipation(ctx context.Context, request *pb.EnableParticipationRequest) (*pb.EnableParticipationResponse, error) {
	err := s.app.EnableParticipation(ctx, commands.EnableParticipation{
		ID: request.GetId(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.EnableParticipationResponse{}, nil
}
func (s server) DisableParticipation(ctx context.Context, request *pb.DisableParticipationRequest) (*pb.DisableParticipationResponse, error) {
	err := s.app.DisableParticipation(ctx, commands.DisableParticipation{
		ID: request.GetId(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.DisableParticipationResponse{}, nil
}

func (s server) storeFromDomain(store *domain.Store) *pb.Store {
	return &pb.Store{
		Id:            store.ID,
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func (s server) productFromDomain(product *domain.Product) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		StoreId:     product.StoreID,
		Name:        product.Name,
		Description: product.Description,
		Sku:         product.SKU,
		Price:       product.Price,
	}

}
