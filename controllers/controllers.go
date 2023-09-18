package controllers

import (
	"my_learnings/render"
	"net/http"

	"github.com/gin-gonic/gin"
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
	if c.Request.Method == "POST" {
		//TODO:
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
	if c.Request.Method == "POST" {
		//TODO:

		//
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Request method not allowed"})
	}
}
func HomePage(c *gin.Context) {
	if c.Request.Method == "GET" || c.Request.Method == "POST" {
		if err := render.RenderTemplate(c, "home"); err != nil {
			// Handle the error, e.g., return an error response
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Request method not allowed"})
	}
}
