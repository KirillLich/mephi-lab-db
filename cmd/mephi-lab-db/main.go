package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type bibki struct {
	gorm.Model
	number int
}

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&bibki{})
}
