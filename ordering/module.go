package ordering

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/monolith"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	return nil
}
