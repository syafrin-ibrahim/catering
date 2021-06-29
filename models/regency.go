package models

import "time"

type Regency struct {
	ID           int           `gorm:"primary_key" json:"id"`
	Name         string        `json:"name"`
	ShippingCost int           `json:"shipping_cost"`
	Transaction  []Transaction `gorm:"ForeignKey:RegencyID;" json:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
