package model

import (
	"time"

	"gorm.io/gorm"
)

type Form struct {
	ID          int64          `gorm:"primaryKey" json:"id"`
	TenantID    int64          `gorm:"not null" json:"tenant_id"`
	FormKey     string         `gorm:"type:varchar(64);uniqueIndex;not null" json:"form_key"`
	Name        string         `gorm:"type:text;not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	CategoryID  *int64         `json:"category_id"`
	Layout      string         `gorm:"type:text;default:single-column" json:"layout"`
	Status      int            `gorm:"default:1" json:"status"`
	Version     int            `gorm:"default:1" json:"version"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Fields      []FormField    `gorm:"foreignKey:FormID" json:"fields,omitempty"`
}
