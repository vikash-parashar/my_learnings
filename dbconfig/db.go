// dbconfig/dbconfig.go

package dbconfig

import (
	"my_learnings/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "user=postgres password=vikash dbname=test host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	AutoMigrate()
	return db, nil
}

// Export this function to access the DB instance from other packages
func GetDB() *gorm.DB {
	return db
}

// AutoMigrate performs auto migration for defined models
func AutoMigrate() {
	db.AutoMigrate(&models.Book{}, &models.Author{}, &models.User{}, &models.Address{}, &models.Order{})
}

// GetHashedPasswordFromDB retrieves the hashed password from the PostgreSQL database based on the user's email.
func GetHashedPasswordFromDB(db *gorm.DB, email string) (string, error) {
	// Implement the SQL query to fetch the hashed password based on the email
	var hashedPassword string
	err := db.Raw("SELECT password FROM users WHERE email = ?", email).Scan(&hashedPassword).Error
	if err != nil {
		return "", err
	}

	return hashedPassword, nil
}
