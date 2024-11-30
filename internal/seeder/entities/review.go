package entities

import (
	"mephi-lab-db/internal/models"
	"mephi-lab-db/internal/seeder/selector"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type ReviewSeeder interface {
	Seed(count uint)
	SetDishIDs(cardIDs []uint)
	SetUserIDs(userIDs []uint)
}

type ReviewSeederImpl struct {
	db      *gorm.DB
	dishIDs []uint
	userIDs []uint
}

func NewReviewSeeder(db *gorm.DB) ReviewSeeder {
	return &ReviewSeederImpl{db: db}
}

func (s *ReviewSeederImpl) Seed(count uint) {
	possibleRating := []uint{0, 1, 2, 3, 4, 5}

	reviews := make([]models.Review, count)
	for i := uint(0); i < count; i++ {
		reviews[i] = models.Review{
			ReviewID: i + 1,
			DishID:   selector.NewSelector().RandomSelect(s.dishIDs),
			UserID:   selector.NewSelector().RandomSelect(s.userIDs),
			Text:     gofakeit.Comment(),
			Rating:   selector.NewSelector().RandomSelect(possibleRating),
		}
	}

	if err := s.db.Create(&reviews).Error; err != nil {
		panic(err)
	}
}

func (s *ReviewSeederImpl) SetDishIDs(ids []uint) {
	s.dishIDs = ids
}

func (s *ReviewSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}
