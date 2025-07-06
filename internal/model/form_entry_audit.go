package model

import (
	"time"
)

type FormEntryAudit struct {
	ID         int64      `gorm:"primaryKey" json:"id"`
	EntryID    int64      `gorm:"index;not null" json:"entry_id"`
	Step       int        `json:"step"`
	ApproverID int64      `json:"approver_id"`
	Status     string     `gorm:"type:text;default:'pending'" json:"status"`
	Comment    string     `json:"comment"`
	ApprovedAt *time.Time `json:"approved_at"`
	CreatedAt  time.Time  `json:"created_at"`
}
