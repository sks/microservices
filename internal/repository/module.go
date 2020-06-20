package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	// postgres connection dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sks/microservices/internal/env"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const localhostPGConnectionString = "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable"

// In ...
type In struct {
	fx.In
	EnvFx     env.Environment
	Logger    *zap.Logger
	LifeCycle fx.Lifecycle
}

// Out ...
type Out struct {
	fx.Out
	DB                  *gorm.DB
	HealthCheckFunction func(ctx context.Context) (interface{}, error) `name:"dbHealthCheckFunction"`
}

// Module ...
func Module(in In) (Out, error) {
	logger := in.Logger.Named("repository")
	db, err := gorm.
		Open("postgres",
			in.EnvFx.GetValOrDefault("PG_CONNECTION_STRING", localhostPGConnectionString))
	if err != nil {
		return Out{}, err
	}
	in.LifeCycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
	db.SetLogger(&dblogger{logger: logger})
	return Out{
		DB: db.Debug(),
		HealthCheckFunction: func(ctx context.Context) (interface{}, error) {
			err := db.Debug().Select("select 1").Error
			if err != nil {
				return nil, err
			}
			return "SUCCESS", nil
		},
	}, nil
}
