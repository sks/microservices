package grpcserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sks/microservices/internal/env"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// GwIn ...
type GwIn struct {
	fx.In
	Lifecycle fx.Lifecycle
	Envfx     env.Environment
	Logger    *zap.Logger
}

// GwOut ...
type GwOut struct {
	fx.Out
	MuxServer          *runtime.ServeMux
	GRPCServerEndpoint string `name:"grpc_server_endpoint"`
}

// GatewayModule create a new gateway server for grpc
func GatewayModule(in GwIn) (GwOut, error) {
	gwPort := in.Envfx.GetValOrDefault("GW_PORT", "7001")
	logger := in.Logger.Named("gateway_server").With(
		zap.String("gateway_port", gwPort))
	muxServer := runtime.NewServeMux()
	httpServer := http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		Addr:              fmt.Sprintf(":%s", gwPort),
		Handler:           muxServer,
	}
	in.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go httpServer.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping the gatewary server")
			return httpServer.Shutdown(ctx)
		},
	})
	return GwOut{
		MuxServer:          muxServer,
		GRPCServerEndpoint: fmt.Sprintf("localhost:%s", getGRPCPort(in.Envfx)),
	}, nil
}
