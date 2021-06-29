package models

import "time"

type Paket struct {
	ID          int           `gorm:"primary_key" json:"id"`
	CategoryID  int           `json:"category_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Transaction []Transaction `gorm:"ForeignKey:PaketID;" json:"-"`
	Price       int           `json:"price"`
	Discount    int           `json:"discount"`
	Image       []Image
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaketResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Discount    int    `json:"discount"`
	Image       string
}
