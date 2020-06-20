package featureflag

import (
	context "context"

	"github.com/sks/microservices/internal/env"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// ClientFxIn dependencies for this module
type ClientFxIn struct {
	fx.In
	Envfx    env.Environment
	Lifecyle fx.Lifecycle
}

// NewClientModule create new server module
func NewClientModule(input ClientFxIn) (ServiceClient, error) {
	conn, err := grpc.Dial(
		input.Envfx.GetValOrDefault("FEATURE_FLAG_SERVER_ADDR", "localhost:7000"))
	if err != nil {
		return nil, err
	}

	input.Lifecyle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})
	return NewServiceClient(conn), nil
}
