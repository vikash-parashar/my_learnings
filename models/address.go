package models

import "gorm.io/gorm"

// Address represents an address model
type Address struct {
	gorm.Model
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}
