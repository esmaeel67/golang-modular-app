package application

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/stores/storespb"
)

type StoreHandlers[T ddd.Event] struct {
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*StoreHandlers[ddd.Event])(nil)

func NewStoreHandlers(logger logger.Logger) StoreHandlers[ddd.Event] {
	return StoreHandlers[ddd.Event]{
		logger: logger,
	}
}

func (h StoreHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case storespb.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storespb.StoreParticipatingToggledEvent:
		return h.onStoreParticipationToggled(ctx, event)
	case storespb.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	}
	return nil
}

func (h StoreHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreCreated)
	message := fmt.Sprintf(`ID: %s, Name: "%s", Location: "%s"`, payload.GetId(), payload.GetName(), payload.GetLocation())
	h.logger.Info(logger.Baskets, logger.StoreCreated, message, nil)
	return nil
}

func (h StoreHandlers[T]) onStoreParticipationToggled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreParticipationToggled)
	message := fmt.Sprintf(`ID: %s, Participating: %b`, payload.GetId(), payload.Participating)
	h.logger.Info(logger.Baskets, logger.StoreParticipationToggled, message, nil)
	return nil
}

func (h StoreHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreRebranded)
	message := fmt.Sprintf(`ID: %s, Name: "%s"`, payload.GetId(), payload.GetName())
	h.logger.Info(logger.Baskets, logger.StoreRebranded, message, nil)
	return nil
}
