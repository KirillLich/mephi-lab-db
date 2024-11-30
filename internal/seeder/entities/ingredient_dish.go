package entities

import (
	"math/rand"
	"mephi-lab-db/internal/models"
	"mephi-lab-db/internal/seeder/selector"

	"gorm.io/gorm"
)

const (
	minIngredientsInDish = 3
	maxIngredientInDish  = 10
)

type IngredientDishSeeder interface {
	Seed(count uint)
	SetIngredientIDs(ingredientIDs []uint)
	SetDishIDs(dishIDs []uint)
}

type IngredientDishSeederImpl struct {
	db             *gorm.DB
	ids            []uint
	ingredientsIDs []uint
	dishIDs        []uint
}

func NewIngredientDishSeeder(db *gorm.DB) IngredientDishSeeder {
	return &IngredientDishSeederImpl{db: db}
}

func (s *IngredientDishSeederImpl) Seed(count uint) {
	internalCount := count

	var connections []models.IngredientDish
	//maxID := 0
	for _, currentDish := range s.dishIDs {

		currentCount := uint(rand.Intn(maxIngredientInDish-minIngredientsInDish) + minIngredientsInDish)
		if currentCount > internalCount {
			currentCount = internalCount
		} else {
			internalCount -= currentCount
		}

		currentConnections := make([]models.IngredientDish, currentCount)

		for i := uint(0); i < currentCount; i++ {
			currentConnections[i] = models.IngredientDish{
				//ID:           uint(i + 1 + uint(maxID)),
				IngredientID: selector.NewSelector().RandomSelect(s.ingredientsIDs),
				DishID:       currentDish,
			}
		}
		connections = append(connections, currentConnections...)
	}

	if err := s.db.Create(&connections).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, len(connections))
	for i, _ := range connections {
		s.ids[i] = uint(i)
	}
}

func (s *IngredientDishSeederImpl) SetDishIDs(ids []uint) {
	s.dishIDs = ids
}

func (s *IngredientDishSeederImpl) SetIngredientIDs(ids []uint) {
	s.ingredientsIDs = ids
}
