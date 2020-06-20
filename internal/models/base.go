package models

import "time"

var now = time.Now

// Base ...
type Base struct {
	CreatedOn time.Time `gorm:"created_on"`
	UpdatedOn time.Time `gorm:"updated_on"`
	CreatedBy string    `gorm:"created_by"`
	UpdatedBy string    `gorm:"updated_by"`
}

// NewBase create a new Base entry
func NewBase(createdBy string) Base {
	return Base{now(), now(), createdBy, createdBy}
}
