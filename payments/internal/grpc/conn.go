package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Dial(ctx context.Context, endpoint string) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.NewClient(endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	// Clean up if context is cancelled
	go func() {
		<-ctx.Done()
		if closeErr := conn.Close(); closeErr != nil {
			// Log error when logging is available
			// log.Printf("failed to close grpc connection: %v", closeErr)
		}
	}()

	return
}
