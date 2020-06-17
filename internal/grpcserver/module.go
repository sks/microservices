package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/sks/microservices/internal/env"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type In struct {
	fx.In
	Lifecycle  fx.Lifecycle
	Logger     *zap.Logger
	Envfx      env.Environment
	Shutdowner fx.Shutdowner
}

type Out struct {
	fx.Out
	GRPCServer *grpc.Server
}

// Module create a new port listener
func Module(in In) (Out, error) {
	port := in.Envfx.GetValOrDefault("PORT", "7000")
	logger := in.Logger.Named("grpcserver").With(
		zap.String("module", "grpcserver"),
		zap.String("port", port))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return Out{}, fmt.Errorf("Error listening to the port: %w", err)
	}
	logger.Info("Started listeing on port")
	grpcServer := grpc.NewServer()
	in.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting the GRPC Server")
				err := grpcServer.Serve(listener)
				if err != nil {
					logger.Error("Error starting the GRPC Server", zap.Error(err))
					in.Shutdowner.Shutdown()
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := listener.Close()
			if err != nil {
				logger.Error("error stopping the listener Server", zap.Error(err))
			}
			grpcServer.GracefulStop()
			logger.Info("Stopped the port listener and GRPC Server")
			return nil
		},
	})
	return Out{
		GRPCServer: grpcServer,
	}, nil
}
