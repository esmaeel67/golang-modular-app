package application

import (
	"context"
	"fmt"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/logger"
	"github.com/esmaeel67/golang-modular-app/stores/storespb"
)

type ProductHandlers[T ddd.Event] struct {
	logger logger.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*ProductHandlers[ddd.Event])(nil)

func NewProductHandlers(logger logger.Logger) ProductHandlers[ddd.Event] {
	return ProductHandlers[ddd.Event]{
		logger: logger,
	}
}

func (h ProductHandlers[T]) HandleEvent(ctx context.Context, event T) error {

	switch event.EventName() {
	case storespb.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case storespb.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case storespb.ProductPriceIncreasedEvent, storespb.ProductPriceDecreasedEvent:
		return h.onProductPriceChanged(ctx, event)
	case storespb.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}
	return nil
}
func (h ProductHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductAdded)

	message := fmt.Sprintf(`ProductHandlers: ID: %s, Name: "%s", Price: "%d"`, payload.GetId(), payload.GetName(), payload.GetPrice())
	h.logger.Info(logger.Baskets, logger.ProductAdded, message, nil)

	return nil
}

func (h ProductHandlers[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRebranded)
	message := fmt.Sprintf(`ID: %s, Name: "%s", Description: "%s"`, payload.GetId(), payload.GetName(), payload.GetDescription())
	h.logger.Info(logger.Baskets, logger.ProductRebranded, message, nil)
	return nil
}

func (h ProductHandlers[T]) onProductPriceChanged(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductPriceChanged)
	message := fmt.Sprintf(`ID: %s, Delta: "%d"`, payload.GetId(), payload.GetDelta())
	h.logger.Info(logger.Baskets, logger.ProductPriceChanged, message, nil)
	return nil
}

func (h ProductHandlers[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRemoved)
	message := fmt.Sprintf(`ID: %s, Price: "%d"`, payload.GetId())
	h.logger.Info(logger.Baskets, logger.ProductRemoved, message, nil)
	return nil
}
