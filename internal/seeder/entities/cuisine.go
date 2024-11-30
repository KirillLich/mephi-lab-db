package entities

import (
	"mephi-lab-db/internal/models"

	"gorm.io/gorm"
)

type CuisineSeeder interface {
	Seed()
	GetIDs() []uint
}

type CuisineSeederImpl struct {
	db  *gorm.DB
	ids []uint
}

func NewCuisineSeeder(db *gorm.DB) CuisineSeeder {
	return &CuisineSeederImpl{db: db}
}

func (s *CuisineSeederImpl) Seed() {

	CuisinesNames := [...]string{"African", "Asian", "American", "British", "Chinese", "European", "French", "German", "Greek", "Indian", "Spanish", "Mexican"}

	cuisines := make([]models.Cuisine, len(CuisinesNames))
	for i, cuisine := range CuisinesNames {
		cuisines[i] = models.Cuisine{
			CuisineID: uint(i + 1),
			Country:   cuisine,
		}
	}

	if err := s.db.Create(&cuisines).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, len(CuisinesNames))
	for i, cuisine := range cuisines {
		s.ids[i] = cuisine.CuisineID
	}
}

func (s *CuisineSeederImpl) GetIDs() []uint {
	return s.ids
}
