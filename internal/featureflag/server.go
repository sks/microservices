package featureflag

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/sks/microservices/internal/models"
	"go.uber.org/zap"
)

type server struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (s *server) GetFlag(ctx context.Context, req *GetFlagRequest) (*GetFlagResponse, error) {
	val, err := s.getVal(req.Key, req.TenantId)
	if err != nil {
		return nil, err
	}
	return &GetFlagResponse{
		Value: val.Value,
	}, err
}

func (s *server) CreateFlag(ctx context.Context, req *CreateFlagRequest) (*CreateFlagResponse, error) {
	val, err := s.getVal(req.Key, req.TenantId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		val = Entity{
			TenantAwareEntity: models.NewTenantAwareEntity(req.TenantId, `//TODO admin??`),
			Key:               req.Key,
		}
	}
	val.Value = req.Value
	val.Description = req.Description
	db := s.db.BeginTx(ctx, &sql.TxOptions{})

	err = db.Save(&val).Error
	if err != nil {
		return nil, err
	}
	return &CreateFlagResponse{}, nil
}

func (s *server) getVal(key, tenantID string) (val Entity, err error) {
	err = s.db.
		Where("key=?", key).
		Where("tenant_id=?", tenantID).
		First(&val).Error
	return
}
