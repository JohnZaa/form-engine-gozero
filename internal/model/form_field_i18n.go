package model

type FormFieldI18n struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	FieldID     int64  `gorm:"index;not null" json:"field_id"`
	LangCode    string `gorm:"type:varchar(8);not null" json:"lang_code"`
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
}
