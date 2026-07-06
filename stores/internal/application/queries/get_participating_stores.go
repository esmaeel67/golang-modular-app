package queries

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
)

type GetParticipatingStores struct {
}

type GetParticipatingStoriesHandler struct {
	mall domain.MallRepository
}

func NewGetParticipatingStoresHandler(mall domain.MallRepository) GetParticipatingStoriesHandler {
	return GetParticipatingStoriesHandler{mall: mall}
}

func (h GetParticipatingStoriesHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*domain.MallStore, error) {
	return h.mall.AllParticipating(ctx)
}
