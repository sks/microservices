package main

import (
	"time"

	"github.com/sks/microservices/internal/env"
	"github.com/sks/microservices/internal/featureflag"
	"github.com/sks/microservices/internal/grpcserver"
	"github.com/sks/microservices/internal/logging"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.StopTimeout(10*time.Second),
		fx.StartTimeout(10*time.Second),
		fx.Provide(
			env.NewEnvFx,
			logging.NewLogger,
			featureflag.NewModule,
			grpcserver.Module,
		),
		fx.Invoke(func(_ featureflag.ServiceServer) {}),
	)
	app.Run()
}
