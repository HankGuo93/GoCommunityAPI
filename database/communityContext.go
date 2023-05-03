package database

import (
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	connectionString := os.Getenv("COMMUNITY_DB_URL")
	defer func() {
		if innerErr := recover(); innerErr != nil {
			for innerErr != nil {
				time.Sleep(5 * time.Second)
				DB, innerErr = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
			}
		}
	}()
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
