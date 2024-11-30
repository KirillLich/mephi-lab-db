package entities

import (
	"mephi-lab-db/internal/models"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type UserSeeder interface {
	Seed(count uint)
	GetIDs() []uint
}

type UserSeederImpl struct {
	db  *gorm.DB
	ids []uint
}

func NewUserSeeder(db *gorm.DB) UserSeeder {
	return &UserSeederImpl{db: db}
}

func (s *UserSeederImpl) Seed(count uint) {

	users := make([]models.User, count)
	for i := uint(0); i < count; i++ {
		users[i] = models.User{
			UserID: i + 1,
			Name:   gofakeit.Username(),
			PWhash: gofakeit.Password(true, true, true, false, false, 7),
		}
	}

	if err := s.db.Create(&users).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, user := range users {
		s.ids[i] = user.UserID
	}
}

func (s *UserSeederImpl) GetIDs() []uint {
	return s.ids
}
