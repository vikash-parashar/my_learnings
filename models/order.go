package models

// Order represents an order model with book and user details
type Order struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	BookID    uint      `json:"book_id"`
	Book      Book      `gorm:"foreignkey:BookID" json:"book"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignkey:UserID" json:"user"`
	OrderDate string    `json:"order_date"`
	OrderType OrderType `json:"order_type"`
}
