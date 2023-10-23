package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
)

var (
	DBConn *gorm.DB
)

// move init and close into database package
func Init() error {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Database connection successfully opened.")
	return nil
}

func Close() {
	sqlDB, err := DBConn.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying database connection: %v", err)
		return
	}
	sqlDB.Close()
}
