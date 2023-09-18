package models

import "gorm.io/gorm"

// Define constants for user and order types (if used throughout your application)
const (
	UserTypeAdmin   UserType = "Admin"
	UserTypeGeneral UserType = "General"
)

// User represents a user model with address details
type User struct {
	gorm.Model
	ID        uint     `gorm:"primary_key" json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	AddressID uint     `json:"address_id"`
	Address   Address  `gorm:"foreignkey:AddressID" json:"address"`
	OTP       string   `json:"otp"`
	UserType  UserType `json:"user_type"`
	Password  []byte   `json:"password"`
}
