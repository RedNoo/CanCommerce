package models

import (
	"time"
)

// ProductColor Model Description
type ProductColor struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Value     int    `json:"value"`
	Quantity  int    `json:"quantity"`
	ProductId int64  `json:"product_id"`

	CreatedAt time.Time  `json:"createdat"`
	DeletedAt *time.Time `json:"deletedat"`
	UpdatedAt *time.Time `json:"updatedat"`
}
