package entities

import (
	"math/rand"
	"mephi-lab-db/internal/models"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type RecipeStepSeeder interface {
	Seed(count uint)
	SetDishIDs(cardIDs []uint)
}

type RecipeStepSeederImpl struct {
	db      *gorm.DB
	ids     []uint
	dishIDs []uint
}

func NewRecipeStepSeeder(db *gorm.DB) RecipeStepSeeder {
	return &RecipeStepSeederImpl{db: db}
}

func (s *RecipeStepSeederImpl) Seed(count uint) {
	internalCount := count
	var recipeStep []models.RecipeStep
	maxID := 0

	for _, dishID := range s.dishIDs {
		stepNumber := uint(rand.Intn(15) + 1)
		if stepNumber > internalCount {
			stepNumber = internalCount
		} else {
			internalCount -= stepNumber
		}

		currentRecipe := make([]models.RecipeStep, stepNumber)
		for i := uint(0); i < stepNumber; i++ {
			currentRecipe[i] = models.RecipeStep{
				RecipeStepID: uint(maxID) + i + 1,
				DishID:       dishID,
				Timing:       uint(rand.Intn(120)),
				Text:         gofakeit.Paragraph(1, 3, 5, " "),
			}
		}
		maxID += int(stepNumber)
		recipeStep = append(recipeStep, currentRecipe...)
	}

	if err := s.db.Create(&recipeStep).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, len(recipeStep))
	for i, step := range recipeStep {
		s.ids[i] = step.RecipeStepID
	}
}

func (s *RecipeStepSeederImpl) SetDishIDs(ids []uint) {
	s.dishIDs = ids
}
