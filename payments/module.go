package payments

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/payments/internal/application"
	"github.com/esmaeel67/golang-modular-app/payments/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/payments/internal/logging"
	"github.com/esmaeel67/golang-modular-app/payments/internal/postgres"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	invoices := postgres.NewInvoiceRepository("invoices", mono.DB())
	payments := postgres.NewPaymentRepository("payments", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	orders := grpc.NewOrderRepository(conn)

	// setup application
	var app application.App
	app = application.New(invoices, payments, orders)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}

	return nil
}
