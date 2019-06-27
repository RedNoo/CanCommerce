package models

import (
	"time"
)

// ProductCategory Model Description
type ProductCategory struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`

	CreatedAt time.Time  `json:"createdat"`
	DeletedAt *time.Time `json:"deletedat"`
	UpdatedAt *time.Time `json:"updatedat"`
}
