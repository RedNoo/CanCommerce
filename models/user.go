package models

import (
	"time"
)

// User Model Description
type User struct {
	ID           int64  `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	UserType     int    `json:"user_type"` //1 admin 2 user
	UserStatusId int    `json:"user_status_id"`

	CreatedAt time.Time  `json:"createdat"`
	DeletedAt *time.Time `json:"deletedat"`
	UpdatedAt *time.Time `json:"updatedat"`
}
