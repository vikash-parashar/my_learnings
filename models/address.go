package models

// Address represents an address model
type Address struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}
