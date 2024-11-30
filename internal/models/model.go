package models

type User struct {
	UserID uint   `gorm:"primarykey"`
	Name   string `gorm:"not null"`
	PWhash string `gorm:"not null"`
}

type Cuisine struct {
	CuisineID uint   `gorm:"primarykey"`
	Country   string `gorm:"not null"`
}

type Dish struct {
	DishID    uint `gorm:"primarykey"`
	CuisineID uint
	Cuisine   Cuisine
	Name      string `json:"title" gorm:"not null"`
	//Type        string `gorm:"not null"`
	EnergyValue uint
	Time        uint
}

type RecipeStep struct {
	RecipeStepID uint `gorm:"primarykey"`
	DishID       uint
	Dish         Dish
	Timing       uint
	Text         string
}

type Favorites struct {
	FavoritesID uint `gorm:"primarykey"`
	UserID      uint `gorm:"unique"`
	User        User
	Dish        []Dish `gorm:"many2many:favorites_dishes"`
}

type Review struct {
	ReviewID uint `gorm:"primarykey"`
	Text     string
	Rating   uint
	UserID   uint
	User     User
	DishID   uint
	Dish     Dish
}

type Article struct {
	ArticleID uint `gorm:"primarykey"`
	UserID    uint
	User      User
	Text      string
	Title     string
	//Date      uint
	Dish []Dish `gorm:"many2many:article_dishes;"`
}

type Ingredient struct {
	IngredientID uint   `json:"id" gorm:"primarykey"`
	Name         string `json:"name"`
	Dish         []Dish `gorm:"many2many:ingredient_dishes;"`
}

type IngredientDish struct {
	IngredientID uint `gorm:"column:ingredient_id"`
	DishID       uint `gorm:"column:dish_id"`
}

type ArticleDish struct {
	ArticleID uint `gorm:"column:article_id"`
	DishID    uint `gorm:"column:dish_id"`
}
