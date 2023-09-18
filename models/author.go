package models

import "gorm.io/gorm"

// Author represents an author model
// Author represents an author model
type Author struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
