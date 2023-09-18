package models

// User represents a user model with address details
type User struct {
	ID        uint     `gorm:"primary_key" json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	AddressID uint     `json:"address_id"`
	Address   Address  `gorm:"foreignkey:AddressID" json:"address"`
	OTP       string   `json:"otp"`
	UserType  UserType `json:"user_type"`
	Password  string   `json:"password"`
}
