package helpers

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"my_learnings/models"
	"net/smtp"
	"time"

	"github.com/twilio/twilio-go"
)

// GenerateOTP generates a random OTP code.
func GenerateOTP(length int) (string, error) {
	if length%2 != 0 {
		return "", errors.New("OTP length must be even")
	}

	// Generate a random byte slice
	randomBytes := make([]byte, length/2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes as a base32 string
	otp := base32.StdEncoding.EncodeToString(randomBytes)
	return otp[:length], nil
}

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

// SendOTPPhone sends an OTP to the provided phone number using Twilio.
func SendOTPPhone(user models.User, otp string) error {
	// Your Twilio Account SID and Auth Token
	accountSID := "your-account-sid"
	authToken := "your-auth-token"

	// Create a Twilio client
	client := twilio.NewRestClient(accountSID, authToken)

	// Replace with your Twilio phone number (must be a purchased Twilio number)
	from := "your-twilio-phone-number"

	// Phone number to which OTP will be sent (in E.164 format, e.g., +1234567890)
	to := user.Phone

	// Create the message body
	message := "Your OTP is: " + otp

	// Send the message
	_, err := client.Messages.SendMessage(from, to, message, nil)
	if err != nil {
		return err
	}

	return nil
}

// VerifyOTP verifies if the provided OTP code matches the stored OTP for the given email/phone.
func VerifyOTP(email, phone, code string) bool {
	var storedOTP models.OTP
	// Assuming you have a database connection using GORM
	if err := models.GetDB().Where("email = ? OR phone = ?", email, phone).Last(&storedOTP).Error; err != nil {
		return false
	}

	// Check if the OTP code matches and is not expired
	if storedOTP.Code == code && storedOTP.ExpiresAt.After(time.Now()) {
		return true
	}
	return false
}
