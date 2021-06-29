package models

import "time"

type User struct {
	ID          int           `gorm:"primary_key" json:"id"`
	Transaction []Transaction `gorm:"ForeignKey:UserID;" json:"-"`
	FullName    string        `json:"full_name"`
	Mobile      string        `json:"mobile"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Address     string        `json:"address"`
	IsAdmin     bool          `json:"is_admin,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserResponse struct {
	UserName string
	Mobile   string
	Email    string
	Token    string
}
