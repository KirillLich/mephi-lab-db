package entities

import (
	"mephi-lab-db/internal/models"

	"gorm.io/gorm"
)

type FavoritesSeeder interface {
	Seed()
	SetUserIDs(userIDs []uint)
}

type FavoritesSeederImpl struct {
	db      *gorm.DB
	ids     []uint
	userIDs []uint
}

func NewFavoritesSeeder(db *gorm.DB) FavoritesSeeder {
	return &FavoritesSeederImpl{db: db}
}

func (s *FavoritesSeederImpl) Seed() {
	favorites := make([]models.Favorites, len(s.userIDs))
	for i, currentID := range s.userIDs {
		favorites[i] = models.Favorites{
			FavoritesID: uint(i + 1),
			UserID:      currentID,
		}
	}

	if err := s.db.Create(&favorites).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, len(s.userIDs))
	for i, favorite := range favorites {
		s.ids[i] = favorite.FavoritesID
	}
}

func (s *FavoritesSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}
