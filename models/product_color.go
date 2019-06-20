package models

import (
	"time"
)

// ProductColor Model Description
type ProductColor struct {
	ID       int64 `gorm:"primary_key"`
	Title    string
	Value    int
	Quantity int
	Product  Product `gorm:"foreignkey:ProductRefer"`

	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m ProductColor) TableName() string {
	return "product_colors"
}
