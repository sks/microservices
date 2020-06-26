package repository

import "go.uber.org/zap"

type dblogger struct {
	logger *zap.Logger
}

func (d *dblogger) Print(v ...interface{}) {
	d.logger.Info("db log", zap.Any("message", v))
}
