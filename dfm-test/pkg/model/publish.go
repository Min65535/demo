package model

import "time"

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	Role      uint      `json:"role"`
}
