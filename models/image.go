package models

import "time"

type Image struct {
	ID        int    `gorm:"primary_key" json:"id"`
	PaketID   int    `form:"paket_id" binding:"required"`
	FileName  string `form:"file" binding:"required"`
	IsMain    bool   `form:"is_main" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ImageResponse struct {
	PaketID  int
	FileName string
	IsMain   bool
}
