package models

import "time"

type Transaction struct {
	ID int `gorm:"primary_key" json:"id"`
	//ID int `gorm:"primary_key" json:"id"`
	//CustomerID int `gorm:"column:customer_id" json:"customer_id" formvalue:"customer_id"`
	UserID     int    `json:"user_id" formvalue:"user_id"`
	PaketID    int    `json:"paket_id" formvalue:"paket_id"`
	Quantity   int    `json:"quantity" formvalue:"quantity"`
	Total      int    `json:"total" formvalue:"total"`
	Location   string `json:"location" formvalue:"location"`
	RegencyID  int    `json:"regency_id" formvalue:"regency_id"`
	Status     string `json:"status" formvalue:"status"`
	PaymentUrl string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TransactionResponse struct {
	UserID     int    `json:"user_id" formvalue:"user_id"`
	PaketID    int    `json:"paket_id" formvalue:"paket_id"`
	Quantity   int    `json:"quantity" formvalue:"quantity"`
	Total      int    `json:"total" formvalue:"total"`
	Status     string `json:"status" formvalue:"status"`
	PaymentUrl string `json:"payment_url" formvalue:"payment"`
}
