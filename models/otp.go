package models

import "time"

// OTP represents a model for storing OTPs.
type OTP struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
}
