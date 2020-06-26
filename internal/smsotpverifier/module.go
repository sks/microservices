package smsotpverifier

import (
	context "context"
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jinzhu/gorm"
	"github.com/sks/microservices/internal/env"
	"go.uber.org/fx"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

type in struct {
	fx.In
	Lifecycle          fx.Lifecycle
	Logger             *zap.Logger
	DB                 *gorm.DB
	GRPCServer         *grpc.Server
	Env                env.Environment
	MuxServer          *runtime.ServeMux
	GRPCServerEndpoint string `name:"grpc_server_endpoint"`
}

type out struct {
	fx.Out
}

// Module create the sms sending module
func Module(i in) (ServiceServer, error) {
	logger := i.Logger.Named("sms_otp_verifier")
	adapter, err := getSMSAdapter(i.Env, logger)
	if err != nil {
		return nil, err
	}
	i.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Migrating DB Schema")
			return i.DB.AutoMigrate(OTPRequest{}).Error
		},
	})
	srvr := &server{
		logger:     logger,
		smsAdapter: adapter,
		db:         i.DB,
	}
	RegisterServiceHandlerFromEndpoint(context.Background(), i.MuxServer,
		i.GRPCServerEndpoint, []grpc.DialOption{grpc.WithInsecure()})
	logger.Debug("Starting GRPC Server")
	RegisterServiceServer(i.GRPCServer, srvr)
	return srvr, nil
}

func getSMSAdapter(e env.Environment, logger *zap.Logger) (Adapter, error) {
	smsProvider, err := e.GetRequiredEnvVar("SMS_PROVIDER")
	if err != nil {
		return nil, err
	}
	switch strings.ToLower(smsProvider) {
	case "d7sms":
		return NewD7SMSProvider(
			e.GetValOrDefault("D7SMS_USERNAME", ""),
			e.GetValOrDefault("D7SMS_PASSWORD", ""),
			logger)
	default:
		return nil, fmt.Errorf("Could not create an SMS adapter for the provider %s", smsProvider)
	}
}
