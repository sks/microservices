package featureflag

import "github.com/sks/microservices/internal/models"

// Entity ...
type Entity struct {
	models.TenantAwareEntity
	Key         string `gorm:"primary_key"`
	Value       string `gorm:"column:value"`
	Description string `gorm:"column:description"`
}

// TableName ...
func (Entity) TableName() string {
	return "feature_flags"
}
