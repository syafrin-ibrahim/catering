package models

import "time"

type Transaction struct {
	ID int `gorm:"primary_key" json:"id"`
	//ID int `gorm:"primary_key" json:"id"`
	//CustomerID int `gorm:"column:customer_id" json:"customer_id" formvalue:"customer_id"`
	UserID      int    `json:"user_id"`
	PaketID     int    `json:"paket_id"`
	Quantity    int    `json:"quantity"`
	Total       int    `json:"total"`
	Location    string `json:"location"`
	RegencyID   int    `json:"regency_id"`
	Status      string `json:"status"`
	PaymentUrl  string
	DeliverTime string `json:"deliver_time"`
	Note        string `json:"note"`
	Paket       Paket
	Regency     Regency
	User        User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MidtransNotification struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

type TransactionResponse struct {
	TransID       int
	CustomerName  string
	PaketName     string
	Quantity      int
	Total         int
	Location      string
	RegentName    string
	DeliveredTime string
	PaymentURL    string
	Note          string
}
