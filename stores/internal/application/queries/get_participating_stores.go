package queries

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type GetParticipatingStores struct {
}

type GetParticipatingStoriesHandler struct {
	participatingStores domain.ParticipatingStoreRepository
}

func NewGetParticipatingStoresHandler(participatingStores domain.ParticipatingStoreRepository) GetParticipatingStoriesHandler {
	return GetParticipatingStoriesHandler{participatingStores: participatingStores}
}

func (h GetParticipatingStoriesHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*domain.Store, error) {
	return h.participatingStores.FindAll(ctx)
}
