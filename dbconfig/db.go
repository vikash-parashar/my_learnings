// dbconfig/dbconfig.go

package dbconfig

import (
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

	return db, nil
}

// Export this function to access the DB instance from other packages
func GetDB() *gorm.DB {
	return db
}
