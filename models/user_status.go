package models

import (
	"time"
)

// UserStatus Model Description
type UserStatus struct {
	ID    int64 `gorm:"primary_key"`
	Title string

	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m UserStatus) TableName() string {
	return "user_statuses"
}
