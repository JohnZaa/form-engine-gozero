package model

import (
	"time"

	"gorm.io/datatypes"
)

type FormVersion struct {
	ID             int64          `gorm:"primaryKey" json:"id"`
	FormID         int64          `gorm:"index;not null" json:"form_id"`
	Version        int            `json:"version"`
	SchemaSnapshot datatypes.JSON `json:"schema_snapshot"` // JSON snapshot of schema
	CreatedAt      time.Time      `json:"created_at"`
}
