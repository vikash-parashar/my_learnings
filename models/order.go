package models

import "gorm.io/gorm"

const (
	OrderTypeBuy  OrderType = "Buy"
	OrderTypeRent OrderType = "Rent"
)

// Define types for order and user types
type (
	OrderType string
	UserType  string
)

// Order represents an order model with book and user details
type Order struct {
	gorm.Model
	BookID    uint      `json:"book_id"`
	Book      Book      `gorm:"foreignkey:BookID" json:"book"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignkey:UserID" json:"user"`
	OrderDate string    `json:"order_date"`
	OrderType OrderType `json:"order_type"`
}
