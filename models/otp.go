package models

import (
	"time"

	"gorm.io/gorm"
)

// OTP represents a model for storing OTPs.
type OTP struct {
	gorm.Model
	ID        uint      `gorm:"primary_key" json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
}
