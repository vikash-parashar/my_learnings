package models

// Author represents an author model
type Author struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
