package models

// TenantAwareEntity ...
type TenantAwareEntity struct {
	Base
	TenantID string `gorm:"primary_key"`
}

// NewTenantAwareEntity ...
func NewTenantAwareEntity(tenantID, createdBy string) TenantAwareEntity {
	return TenantAwareEntity{
		Base:     NewBase(createdBy),
		TenantID: tenantID,
	}
}
