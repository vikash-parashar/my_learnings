package helpers

import (
	"fmt"
	"my_learnings/dbconfig"
	"my_learnings/models"
	"net/smtp"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/sfreiberg/gotwilio"
)

var (
	db, _ = dbconfig.InitDB()
)

// SendOTPEmail sends an OTP to the provided email address.
func SendOTPEmail(user models.User, otp string) error {
	// Replace the following placeholders with your SMTP server and credentials.
	smtpServer := "your-smtp-server.com"
	smtpPort := 587
	smtpUsername := "your-smtp-username"
	smtpPassword := "your-smtp-password"

	// Set up the SMTP client
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	to := []string{user.Email}
	msg := []byte("Subject: Your OTP\n" +
		"To: " + user.Email + "\n" +
		"\n" +
		"Your OTP is: " + otp)

	// Send the email
	err := smtp.SendMail(smtpServer+":"+fmt.Sprint(smtpPort), auth, smtpUsername, to, msg)
	if err != nil {
		return err
	}

	return nil
}

// SendOTPPhone sends an OTP to the provided phone number.
func SendOTPPhone(u models.User, otp string) error {
	// Generate a new TOTP key
	_, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "YourApp",
		AccountName: u.Phone,
	})
	if err != nil {
		return err
	}

	// Get the OTP URL
	// otpURL := key.URL()

	// Use the OTP URL to send the OTP to the user's phone number using Twilio
	twilio := gotwilio.NewTwilioClient("your-account-sid", "your-auth-token")
	message := "Your OTP is: " + otp
	_, _, err = twilio.SendSMS("your-twilio-phone-number", u.Phone, message, "", "")
	if err != nil {
		return err
	}

	return nil
}

// VerifyOTP verifies if the provided OTP code matches the stored OTP for the given email/phone.
func VerifyOTP(email, phone, code string) bool {
	var storedOTP models.OTP
	// Assuming you have a database connection using GORM
	if err := db.Where("email = ? OR phone = ?", email, phone).Last(&storedOTP).Error; err != nil {
		return false
	}

	// Check if the OTP code matches and is not expired
	if storedOTP.Code == code && storedOTP.ExpiresAt.After(time.Now()) {
		return true
	}
	return false
}

// GenerateOTP generates a new OTP and returns the OTP URL.
func GenerateOTP() (string, error) {
	// Generate a new TOTP key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "YourApp",
		AccountName: "user@example.com",
	})
	if err != nil {
		return "", err
	}

	// Get the OTP URL
	otpURL := key.URL()

	return otpURL, nil
}

// StoreOTP stores an OTP in the database.
func StoreOTP(email, phone, code string, expiration time.Time) error {
	// Create an OTP model
	otpModel := models.OTP{
		Email:     email,
		Phone:     phone,
		Code:      code,
		ExpiresAt: expiration,
	}

	// Store the OTP in the database using GORM
	if err := db.Create(&otpModel).Error; err != nil {
		return err
	}

	return nil
}
