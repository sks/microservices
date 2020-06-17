package featureflag

import (
	"context"

	"go.uber.org/zap"
)

type server struct {
	logger *zap.Logger
}

func (s *server) GetFlag(ctx context.Context, req *GetFlagRequest) (*GetFlagResponse, error) {
	return &GetFlagResponse{}, nil
}
