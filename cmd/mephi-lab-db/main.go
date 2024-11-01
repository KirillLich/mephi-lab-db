package main

import (
	"log"
	"os"

	"mephi-lab-db/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dst := []interface{}{
		&models.User{},
		&models.Cuisine{},
		&models.Dish{},
		&models.RecipeStep{},
		&models.Review{},
		&models.Favorites{},
		&models.Article{},
		&models.Ingredient{},
	}

	db.AutoMigrate(dst...)
}
