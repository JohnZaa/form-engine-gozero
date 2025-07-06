package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FormField struct {
	ID                int64          `gorm:"primaryKey" json:"id"`
	FormID            int64          `gorm:"index;not null" json:"form_id"`
	GroupKey          string         `json:"group_key"`
	GroupLabel        string         `json:"group_label"`
	FieldKey          string         `gorm:"not null" json:"field_key"`
	Label             string         `gorm:"not null" json:"label"`
	Type              string         `gorm:"not null" json:"type"`
	Required          bool           `gorm:"default:false" json:"required"`
	Options           datatypes.JSON `json:"options"`       // 可选项
	DefaultValue      datatypes.JSON `json:"default_value"` // 默认值
	Placeholder       string         `json:"placeholder"`
	Validation        datatypes.JSON `json:"validation"` // 校验规则
	VisibleCondition  datatypes.JSON `json:"visible_condition"`
	EditableCondition datatypes.JSON `json:"editable_condition"`
	OrderNum          int            `gorm:"default:0" json:"order_num"`
	Encrypted         bool           `gorm:"default:false" json:"encrypted"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}
