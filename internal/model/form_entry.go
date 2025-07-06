package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FormEntry struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	FormID    int64          `gorm:"index;not null" json:"form_id"`
	TenantID  int64          `gorm:"index;not null" json:"tenant_id"`
	UserID    int64          `json:"user_id"`
	Data      datatypes.JSON `gorm:"type:jsonb;not null" json:"data"`
	Status    string         `gorm:"type:text;default:'draft'" json:"status"`
	Version   int            `gorm:"default:1" json:"version"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
