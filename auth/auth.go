package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// SecretKey is your JWT secret key for signing and validating tokens.
var SecretKey = []byte("your-secret-key")

// GenerateJWTToken generates a JWT token for the given email.
func GenerateJWTToken(email string) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	// Sign the token with your secret key
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWTToken validates a JWT token and returns the claims if valid.
func ValidateJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// Replace this with your actual JWT token validation logic
func IsValidToken(token string) bool {
	// Implement your JWT validation logic here
	// Return true if the token is valid, false otherwise
	// You may want to use a JWT library to validate the token
	return true
}

// ComparePasswords securely compares the provided password with the hashed password from the database.
func ComparePasswords(providedPassword string, hashedPasswordFromDB string) bool {
	// Use bcrypt.CompareHashAndPassword to securely compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswordFromDB), []byte(providedPassword))
	return err == nil
}
