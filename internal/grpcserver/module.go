package grpcserver

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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

// Out ...
type Out struct {
	fx.Out
	GRPCServer         *grpc.Server
	MuxServer          *runtime.ServeMux
	GRPCServerEndpoint string `name:"grpc_server_endpoint"`
}

// Module create a new port listener
func Module(in In) (Out, error) {
	port := in.Envfx.GetValOrDefault("PORT", "7000")
	gwPort := in.Envfx.GetValOrDefault("GW_PORT", "7001")
	logger := in.Logger.Named("grpcserver").With(
		zap.String("port", port),
		zap.String("gateway_port", gwPort))
	grpcServer := grpc.NewServer()
	muxServer := runtime.NewServeMux()
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
			go http.ListenAndServe(fmt.Sprintf(":%s", gwPort), muxServer)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			logger.Info("Stopped the port listener and GRPC Server")
			return nil
		},
	})
	return Out{
		GRPCServer:         grpcServer,
		MuxServer:          muxServer,
		GRPCServerEndpoint: fmt.Sprintf("localhost:%s", port),
	}, nil
}
