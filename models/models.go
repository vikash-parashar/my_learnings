package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Define constants for user and order types (if used throughout your application)
const (
	UserTypeAdmin   UserType  = "Admin"
	UserTypeGeneral UserType  = "General"
	OrderTypeBuy    OrderType = "Buy"
	OrderTypeRent   OrderType = "Rent"
)

// Define types for order and user types
type (
	OrderType string
	UserType  string
)

var db *gorm.DB
var err error

// InitDB initializes the database connection
func InitDB(connectionString string) (*gorm.DB, error) {
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}

// AutoMigrate performs auto migration for defined models
func AutoMigrate() {
	db.AutoMigrate(&Book{}, &Author{}, &User{}, &Address{}, &Order{})
}
