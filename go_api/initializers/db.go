package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort)
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to conncet with database")
	}

}
