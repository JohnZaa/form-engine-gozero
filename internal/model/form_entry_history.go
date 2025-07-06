package model

import (
	"time"

	"gorm.io/datatypes"
)

type FormEntryHistory struct {
	ID         int64          `gorm:"primaryKey" json:"id"`
	EntryID    int64          `gorm:"index;not null" json:"entry_id"`
	FormID     int64          `gorm:"index;not null" json:"form_id"`
	ChangedBy  int64          `json:"changed_by"`
	OldData    datatypes.JSON `json:"old_data"`
	NewData    datatypes.JSON `json:"new_data"`
	ChangeType string         `gorm:"type:text" json:"change_type"` // create/update/delete
	CreatedAt  time.Time      `json:"created_at"`
}
