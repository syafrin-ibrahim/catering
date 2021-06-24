package models

import "time"

type User struct {
	ID          int           `gorm:"primary_key" json:"id"`
	Transaction []Transaction `gorm:"ForeignKey:UserID;" json:"-"`
	FullName    string        `json:"full_name" formvalue:"full_name"`
	Mobile      string        `json:"mobile" formvalue:"mobile"`
	Email       string        `json:"email" formvalue:"email"`
	Password    string        `gorm:"type:char(60)" json:"password" formvalue:"password"`
	Address     string        `json:"address" formvalue:"address"`
	IsAdmin     bool          `json:"is_admin,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Token    string `json:"token" form:"token"`
}
