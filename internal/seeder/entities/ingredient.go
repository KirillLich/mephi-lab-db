package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mephi-lab-db/internal/models"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

const apiKey = "4badb167e113443393256055ba0b86d4"

type ingredientRespond struct {
	Results      []models.Ingredient `json:"results"`
	Offset       uint                `json:"offset"`
	Number       uint                `json:"number"`
	TotalResults uint                `json:"totalResults"`
}

type IngredientSeeder interface {
	Seed(count uint)
	DummySeed(count uint)
	GetIDs() []uint
}

type IngredientSeederImpl struct {
	db  *gorm.DB
	ids []uint
}

func NewIngredientSeeder(db *gorm.DB) IngredientSeeder {
	return &IngredientSeederImpl{db: db}
}

func (s *IngredientSeederImpl) DummySeed(count uint) {
	type DummyIngredient struct {
		Name string
	}

	type DummyRespond struct {
		Results []DummyIngredient
	}

	var Dummy DummyRespond
	Dummy.Results = make([]DummyIngredient, count/4)
	for i, _ := range Dummy.Results {
		Dummy.Results[i] = DummyIngredient{
			Name: "Ingredient " + gofakeit.Noun(),
		}
	}

	keyWords := []string{"fruits", "vegetables", "nuts", "grains", "meat", "fish", "dairy", "oil", "seasoning", "sauce"}

	var apiCounting uint = 0
	ingredients := make([]models.Ingredient, count)

	for t, _ := range keyWords {
		for _, currentIngredient := range Dummy.Results {
			if apiCounting == count {
				break
			}
			ingredients[apiCounting] = models.Ingredient{
				IngredientID: apiCounting + 1,
				Name:         currentIngredient.Name + " " + keyWords[apiCounting/50+uint(t)],
			}
			apiCounting++
		}
	}

	if err := s.db.Create(&ingredients).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, ingredient := range ingredients {
		s.ids[i] = ingredient.IngredientID
	}
}

func (s *IngredientSeederImpl) Seed(count uint) {
	keyWords := []string{"fruits", "vegetables", "nuts", "grains", "meat", "fish", "dairy", "oil", "seasoning", "sauce"}

	var apiCounting uint = 0
	ingredients := make([]models.Ingredient, count)

	for _, keyWord := range keyWords {
		url := fmt.Sprintf("https://api.spoonacular.com/food/ingredients/search?apiKey=%s&query=%s&number=%d", apiKey, keyWord, 50)
		resp, err := http.Get(url)
		if err != nil {
			// we will get an error at this stage if the request fails, such as if the
			// requested URL is not found, or if the server is not reachable.
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// if we want to check for a specific status code, we can do so here
		// for example, a successful request should return a 200 OK status
		if resp.StatusCode != http.StatusOK {
			// if the status code is not 200, we should log the status code and the
			// status string, then exit with a fatal error
			log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseIngredient ingredientRespond
		err = json.Unmarshal(data, &responseIngredient)
		if err != nil {
			log.Fatal(err)
		}

		for _, currentIngredient := range responseIngredient.Results {
			if apiCounting == count {
				break
			}
			ingredients[apiCounting] = models.Ingredient{
				IngredientID: apiCounting + 1,
				Name:         currentIngredient.Name,
			}
			apiCounting++
		}
	}

	if err := s.db.Create(&ingredients).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, ingredient := range ingredients {
		s.ids[i] = ingredient.IngredientID
	}
}

func (s *IngredientSeederImpl) GetIDs() []uint {
	return s.ids
}
