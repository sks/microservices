package featureflag

import (
	"github.com/sks/microservices/internal/env"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// FxIn dependencies for this module
type FxIn struct {
	fx.In
	Logger     *zap.Logger
	Envfx      env.Environment
	GRPCServer *grpc.Server
}

// Out module outout
type Out struct {
	fx.Out
	ServiceServer ServiceServer
}

// NewModule create new module
func NewModule(input FxIn) (Out, error) {
	logger := input.Logger.Named("featureflag")
	srvr := &server{
		logger,
	}
	logger.Info("Registering the grpc server")
	RegisterServiceServer(input.GRPCServer, srvr)
	return Out{
		ServiceServer: srvr,
	}, nil
}
