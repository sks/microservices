package featureflag

import (
	context "context"

	"github.com/jinzhu/gorm"
	"github.com/sks/microservices/internal/env"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ServerFxIn dependencies for this module
type ServerFxIn struct {
	fx.In
	Logger     *zap.Logger
	Envfx      env.Environment
	GRPCServer *grpc.Server
	Lifecyle   fx.Lifecycle
	DB         *gorm.DB
}

// NewServerModule create new server module
func NewServerModule(input ServerFxIn) (ServiceServer, error) {
	logger := input.Logger.Named("featureflag")
	input.Lifecyle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return input.DB.AutoMigrate(Entity{}).Error
		},
	})
	srvr := &server{
		logger,
		input.DB,
	}
	logger.Info("Registering the grpc server")
	RegisterServiceServer(input.GRPCServer, srvr)
	return srvr, nil
}
