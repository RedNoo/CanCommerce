package models

import (
	"time"
)

// Product Model Description
type Product struct {
	ID               int64 `gorm:"primary_key"`
	SKU              string
	Title            string
	OldPrice         float64 `json:"old_price" sql:"DECIMAL(10,2)"`
	Price            float64 `json:"price" sql:"DECIMAL(10,2)"`
	ShortDescription string
	LongDescription  string
	Tag              string
	Quantity         int
	IsNew            bool
	IsSale           bool
	IsActive         bool

	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Product) TableName() string {
	return "products"
}
