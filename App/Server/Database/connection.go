package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=user password=secret dbname=ciucap port=5432"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}
