// controllers/user_controller.go

package controllers

import (
	"encoding/json"
	"my_learnings/helpers"
	"my_learnings/models"
	"net/http"
	"time"
)

// CreateUser creates a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the request body into a user object
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate OTP codes
	emailOTP, err := helpers.GenerateOTP(6) // Change the length as needed
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	phoneOTP, err := helpers.GenerateOTP(6) // Change the length as needed
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store OTPs in the database
	otpModel := models.OTP{
		Email:     user.Email,
		Phone:     user.Phone,
		Code:      emailOTP,                         // Store email OTP
		ExpiresAt: time.Now().Add(time.Minute * 15), // Adjust expiration time as needed
	}
	if err := models.GetDB().Create(&otpModel).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	otpModel.Code = phoneOTP // Store phone OTP
	if err := models.GetDB().Create(&otpModel).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send OTPs to email and phone
	if err := helpers.SendOTPEmail(user, emailOTP); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := helpers.SendOTPPhone(user, phoneOTP); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Continue with user creation and database operations as before
	// ...

	// Serialize the created user to JSON and send the response
	json.NewEncoder(w).Encode(user)
}
