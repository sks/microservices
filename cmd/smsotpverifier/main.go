package main

import (
	"context"
	"time"

	"github.com/sks/microservices/internal/env"
	"github.com/sks/microservices/internal/grpcserver"
	"github.com/sks/microservices/internal/health"
	"github.com/sks/microservices/internal/logging"
	"github.com/sks/microservices/internal/repository"
	"github.com/sks/microservices/internal/smsotpverifier"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.StopTimeout(10*time.Second),
		fx.StartTimeout(10*time.Second),
		fx.Provide(
			env.NewEnvFx,
			logging.NewLogger,
			smsotpverifier.Module,
			grpcserver.Module,
			repository.Module,
			health.Module,
		),
		fx.Invoke(func(_ smsotpverifier.ServiceServer) {}, registerHealthChecks),
	)
	app.Run()
}

type in struct {
	fx.In
	Reporter              health.Reporter
	DBHealthCheckFunction func(context.Context) (interface{}, error) `name:"dbHealthCheckFunction"`
}

func registerHealthChecks(i in) {
	i.Reporter.AddModule("database", i.DBHealthCheckFunction)
}
