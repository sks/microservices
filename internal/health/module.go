package health

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sks/microservices/internal/env"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// In ...
type In struct {
	fx.In
	Logger    *zap.Logger
	Lifecycle fx.Lifecycle
	Env       env.Environment
}

// Out ...
type Out struct {
	fx.Out
	Reporter Reporter
}

// Module ...
func Module(in In) (Out, error) {
	reporter := newReporter(in.Logger)
	port := in.Env.GetValOrDefault("HEALTH_PORT", "9999")
	healthServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      reporter.server,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	in.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			reporter.logger.
				Info("Starting the health end point",
					zap.String("port", port),
					zap.Int("health_checks", len(reporter.functions)))
			go func() {
				err := healthServer.ListenAndServe()
				if err != nil {
					reporter.logger.Error("Error starting server ", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return healthServer.Shutdown(ctx)
		},
	})
	return Out{
		Reporter: reporter,
	}, nil
}
