package models

import (
	"time"
)

// Product Model Description
type Product struct {
	ID               int64   `json:"id"`
	SKU              string  `json:"sku"`
	Title            string  `json:"title"`
	OldPrice         float64 `json:"old_price"`
	Price            float64 `json:"price"`
	ShortDescription string  `json:"short_description"`
	LongDescription  string  `json:"long_description"`
	Tag              string  `json:"tag"`
	Quantity         int     `json:"quantity"`
	IsNew            bool    `json:"is_new"`
	IsSale           bool    `json:"is_sale"`
	IsActive         bool    `json:"is_active"`

	CreatedAt time.Time  `json:"createdat"`
	DeletedAt *time.Time `json:"deletedat"`
	UpdatedAt *time.Time `json:"updatedat"`
}
