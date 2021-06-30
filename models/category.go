package models

import "time"

type Category struct {
	ID        int     `gorm:"primary_key" json:"id"`
	Name      string  `json:"name"`
	Paket     []Paket `gorm:"ForeignKey:CategoryID;" json:"paket"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryResponse struct {
	ID    int
	Name  string `json:"name"`
	Paket []Paket
}
