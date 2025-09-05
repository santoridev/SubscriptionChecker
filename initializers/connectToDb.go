package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	info := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(info), &gorm.Config{})

	if err != nil {
		log.Fatal("Error conecting to DB")
	}
}
