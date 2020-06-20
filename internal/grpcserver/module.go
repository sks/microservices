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

// In ...
type In struct {
	fx.In
	Lifecycle  fx.Lifecycle
	Logger     *zap.Logger
	Envfx      env.Environment
	Shutdowner fx.Shutdowner
}

// Module create a new port listener
func Module(in In) (*grpc.Server, error) {
	port := in.Envfx.GetValOrDefault("PORT", "7000")
	logger := in.Logger.Named("grpcserver").With(
		zap.String("module", "grpcserver"),
		zap.String("port", port))
	grpcServer := grpc.NewServer()
	in.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
			if err != nil {
				return fmt.Errorf("Error listening to the port: %w", err)
			}
			logger.Info("Started listeing on port")
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
			grpcServer.GracefulStop()
			logger.Info("Stopped the port listener and GRPC Server")
			return nil
		},
	})
	return grpcServer, nil
}
