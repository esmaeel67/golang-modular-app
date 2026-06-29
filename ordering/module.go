package ordering

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/monolith"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/application"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/grpc"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/logging"
	"github.com/esmaeel67/golang-modular-app/ordering/internal/postgres"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	orders := postgres.NewOrderRepository("orders", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	shopping := grpc.NewShoppingRepository(conn)
	notifications := grpc.NewNotificationRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, invoices, shopping, notifications)
	app = logging.NewApplication(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	//TODO: http

	//TODO: swagger

	return nil
}
