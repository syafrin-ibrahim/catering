package models

import "time"

type Paket struct {
	ID          int           `gorm:"primary_key" json:"id"`
	Transaction []Transaction `gorm:"ForeignKey:PaketID;" json:"-"`
	Name        string        `json:"name" formvalue:"name"`
	Description string        `json:"description" formvalue:"description"`
	Price       int           `json:"price" formvalue:"price"`
	Type        string        `json:"type" formvalue:"type"`
	PicturePath string        `json:"picture_path" formvalue:"picture_path"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
