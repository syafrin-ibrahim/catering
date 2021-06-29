package models

import "time"

type Category struct {
	ID        int     `gorm:"primary_key" json:"id"`
	Name      string  `json:"name"`
	Paket     []Paket `gorm:"ForeignKey:CategoryID;" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
