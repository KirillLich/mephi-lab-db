package seeder

import (
	"mephi-lab-db/internal/models"
	"mephi-lab-db/internal/seeder/entities"

	"gorm.io/gorm"
)

const (
	userCount            = 10000
	ingredientCount      = 200
	articleCount         = 700
	dishCount            = 500
	favoritesCount       = 10000
	recipesStepCount     = 15 * dishCount
	reviewCount          = 250
	ingredient_dishCount = 10 * ingredientCount
	articleDishCount     = 10 * dishCount
)

type Seeder interface {
	Seed()
}

type SeederImpl struct {
	db                   *gorm.DB
	userSeeder           entities.UserSeeder
	ingredientSeeder     entities.IngredientSeeder
	articleSeeder        entities.ArticleSeeder
	cuisineSeeder        entities.CuisineSeeder
	dishSeeder           entities.DishSeeder
	favoritesSeeder      entities.FavoritesSeeder
	recipesStepSeeder    entities.RecipeStepSeeder
	reviewSeeder         entities.ReviewSeeder
	ingredientDishSeeder entities.IngredientDishSeeder
	articleDishSeeder    entities.ArticleDishSeeder

	userCount        int
	ingredientCount  int
	articleCount     int
	cuisineCount     int
	dishCount        int
	favoritesCount   int
	recipesStepCount int
	reviewCount      int
}

func NewSeeder(db *gorm.DB) Seeder {
	return &SeederImpl{
		db:                   db,
		userSeeder:           entities.NewUserSeeder(db),
		ingredientSeeder:     entities.NewIngredientSeeder(db),
		articleSeeder:        entities.NewArticleSeeder(db),
		cuisineSeeder:        entities.NewCuisineSeeder(db),
		dishSeeder:           entities.NewDishSeeder(db),
		favoritesSeeder:      entities.NewFavoritesSeeder(db),
		recipesStepSeeder:    entities.NewRecipeStepSeeder(db),
		reviewSeeder:         entities.NewReviewSeeder(db),
		ingredientDishSeeder: entities.NewIngredientDishSeeder(db),
		articleDishSeeder:    entities.NewArticleDishSeeder(db),
	}
}

func (s *SeederImpl) Seed() {
	s.prepare()

	s.seedUsers()
	s.seedIngredients()

	s.seedCuisines()
	s.seedArticles()
	s.seedDishes()

	s.seedIngredientDishConnections()
	s.seedFavorites()
	s.seedRecipeSteps()
	s.seedReviews()
	s.seedArticleDishConnections()
}

func (s *SeederImpl) seedUsers() {
	s.userSeeder.Seed(userCount)
	s.userCount = int(userCount)
}

func (s *SeederImpl) seedIngredients() {
	//s.ingredientSeeder.Seed(ingredientCount)
	s.ingredientSeeder.DummySeed(ingredientCount)
	s.ingredientCount = int(ingredientCount)
}

func (s *SeederImpl) seedArticles() {
	s.articleSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.articleSeeder.Seed(articleCount)
	s.articleCount = int(articleCount)
}

func (s *SeederImpl) seedCuisines() {
	s.cuisineSeeder.Seed()
	s.cuisineCount = 12
}

func (s *SeederImpl) seedDishes() {
	s.dishSeeder.SetCuisineIDs(s.cuisineSeeder.GetIDs())
	//s.dishSeeder.Seed(dishCount)
	s.dishSeeder.DummySeed(dishCount)
	s.dishCount = int(dishCount)
}

func (s *SeederImpl) seedFavorites() {
	s.favoritesSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.favoritesSeeder.Seed()
	s.favoritesCount = int(favoritesCount)
}

func (s *SeederImpl) seedRecipeSteps() {
	s.recipesStepSeeder.SetDishIDs(s.dishSeeder.GetIDs())
	s.recipesStepSeeder.Seed(recipesStepCount)
	s.recipesStepCount = int(recipesStepCount)
}

func (s *SeederImpl) seedReviews() {
	s.reviewSeeder.SetDishIDs(s.dishSeeder.GetIDs())
	s.reviewSeeder.SetUserIDs(s.userSeeder.GetIDs())
	s.reviewSeeder.Seed(reviewCount)
	s.reviewCount = int(reviewCount)
}

func (s *SeederImpl) seedIngredientDishConnections() {
	s.ingredientDishSeeder.SetIngredientIDs(s.ingredientSeeder.GetIDs())
	s.ingredientDishSeeder.SetDishIDs(s.dishSeeder.GetIDs())
	s.ingredientDishSeeder.Seed(ingredient_dishCount)
	s.reviewCount = int(reviewCount)
}

func (s *SeederImpl) seedArticleDishConnections() {
	s.articleDishSeeder.SetArticleIDs(s.articleSeeder.GetIDs())
	s.articleDishSeeder.SetDishIDs(s.dishSeeder.GetIDs())
	s.articleDishSeeder.Seed(articleDishCount)
	s.articleCount = int(articleCount)
}

func (s *SeederImpl) prepare() {
	s.dropAll()
	s.migrateAll()
}

func (s *SeederImpl) dropAll() {
	tables, err := s.db.Migrator().GetTables()
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		if err := s.db.Migrator().DropTable(table); err != nil {
			panic(err)
		}
	}
}

func (s *SeederImpl) migrateAll() {
	if err := s.db.AutoMigrate(
		&models.User{},
		&models.Ingredient{},
		&models.Article{},
		&models.Cuisine{},
		&models.Dish{},
		&models.Favorites{},
		&models.RecipeStep{},
		&models.Review{},
		&models.IngredientDish{},
		&models.ArticleDish{},
	); err != nil {
		panic(err)
	}
}
