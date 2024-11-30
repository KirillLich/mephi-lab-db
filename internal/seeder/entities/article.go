package entities

import (
	"mephi-lab-db/internal/models"
	"mephi-lab-db/internal/seeder/selector"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type ArticleSeeder interface {
	Seed(count uint)
	SetUserIDs(userIDs []uint)
	GetIDs() []uint
}

type ArticleSeederImpl struct {
	db      *gorm.DB
	ids     []uint
	dishIDs []uint
	userIDs []uint
}

func NewArticleSeeder(db *gorm.DB) ArticleSeeder {
	return &ArticleSeederImpl{db: db}
}

func (s *ArticleSeederImpl) Seed(count uint) {

	articles := make([]models.Article, count)
	for i := uint(0); i < count; i++ {
		articles[i] = models.Article{
			ArticleID: i + 1,
			UserID:    selector.NewSelector().RandomSelect(s.userIDs),
			Text:      gofakeit.Paragraph(4, 5, 20, " "),
			Title:     cases.Title(language.English, cases.Compact).String(gofakeit.Adjective() + " " + gofakeit.Noun()),
			//Date:      0,
		}
	}

	if err := s.db.Create(&articles).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, article := range articles {
		s.ids[i] = article.ArticleID
	}
}

func (s *ArticleSeederImpl) SetUserIDs(ids []uint) {
	s.userIDs = ids
}

func (s *ArticleSeederImpl) GetIDs() []uint {
	return s.ids
}
