package controllers

import (
	"fmt"
	"my_learnings/auth"
	"my_learnings/dbconfig"
	"my_learnings/models"
	"my_learnings/render"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	db, _ = dbconfig.InitDB()
)

func UserProfile(c *gin.Context) {
	// TODO: Implement logic to retrieve and display the user's profile
}

func RegistrationPage(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		if err := render.RenderTemplate(c, "signup"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	} else if c.Request.Method == http.MethodPost {
		// Handle the POST request for form submissions
		// TODO: Implement registration logic
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Request method not allowed"})
	}
}

func Registration(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		var user models.User // Assuming you have defined the User model

		// Parse user data from the request
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Perform validation for all fields (you can add more checks)
		if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Phone == "" || user.AddressID == 0 || len(user.Password) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}

		// Hash the user's password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
			return
		}

		// Store the hashed password in the User struct
		user.Password = hashedPassword

		// Save the user data to the database using GORM
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user data"})
			return
		}

		// Send a success response or perform any other necessary actions
		c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Request method not allowed"})
	}
}

func LoginPage(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		if err := render.RenderTemplate(c, "login"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	} else if c.Request.Method == http.MethodPost {
		// Handle the POST request for login form submissions
		// TODO: Implement login logic
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Request method not allowed"})
	}
}

func Login(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		// Parse user credentials from the request (e.g., email and password)
		var userCredentials models.User
		if err := c.ShouldBindJSON(&userCredentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Retrieve hashed password from the PostgreSQL database based on the user's email
		hashedPasswordFromDB, err := dbconfig.GetHashedPasswordFromDB(db, userCredentials.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user data"})
			return
		}

		// Compare the provided password with the hashed password from the database
		if auth.ComparePasswords(string(userCredentials.Password), hashedPasswordFromDB) {
			// Passwords match, generate a JWT token
			token, err := auth.GenerateJWTToken(userCredentials.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT token"})
				return
			}

			// Send the JWT token in the response
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		} else {
			// Passwords do not match, send an "Unauthorized" response
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Request method not allowed"})
		return
	}
}

func HomePage(c *gin.Context) {
	// Check if the request has a JWT token in the request body
	token := c.GetHeader("Authorization")

	if token != "" {
		// Validate the JWT token here (you should implement your validation logic)
		if auth.IsValidToken(token) {
			// Token is valid, render the "home" template
			if err := render.RenderTemplate(c, "home"); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			return
		} else {
			// Token is invalid, send a "Not Authorized" response
			fmt.Println("------------------------------------------------")
			fmt.Println("                 Unauthorized ")
			fmt.Println("         message : jwt token is invalid ")
			fmt.Println("________________________________________________")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			return
		}
	} else {
		fmt.Println("------------------------------------------------")
		fmt.Println("                 Unauthorized ")
		fmt.Println("         message : user is not logged in ")
		fmt.Println("________________________________________________")
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
}
