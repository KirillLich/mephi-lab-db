package entities

import (
	"math/rand"
	"mephi-lab-db/internal/models"
	"mephi-lab-db/internal/seeder/selector"

	"gorm.io/gorm"
)

const (
	minDishesInArticle = 3
	maxDishesInArticle = 10
)

type ArticleDishSeeder interface {
	Seed(count uint)
	SetArticleIDs(ingredientIDs []uint)
	SetDishIDs(dishIDs []uint)
}

type ArticleDishSeederImpl struct {
	db         *gorm.DB
	ids        []uint
	articleIDs []uint
	dishIDs    []uint
}

func NewArticleDishSeeder(db *gorm.DB) ArticleDishSeeder {
	return &ArticleDishSeederImpl{db: db}
}

func (s *ArticleDishSeederImpl) Seed(count uint) {
	internalCount := count

	var connections []models.ArticleDish
	//maxID := 0
	for _, currentArticle := range s.articleIDs {

		currentCount := uint(rand.Intn(maxDishesInArticle-minDishesInArticle) + minDishesInArticle)
		if currentCount > internalCount {
			currentCount = internalCount
		} else {
			internalCount -= currentCount
		}

		currentConnections := make([]models.ArticleDish, currentCount)

		for i := uint(0); i < currentCount; i++ {
			currentConnections[i] = models.ArticleDish{
				//ID:           uint(i + 1 + uint(maxID)),
				ArticleID: currentArticle,
				DishID:    selector.NewSelector().RandomSelect(s.dishIDs),
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

func (s *ArticleDishSeederImpl) SetDishIDs(ids []uint) {
	s.dishIDs = ids
}

func (s *ArticleDishSeederImpl) SetArticleIDs(ids []uint) {
	s.articleIDs = ids
}
