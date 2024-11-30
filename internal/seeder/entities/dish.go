package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mephi-lab-db/internal/models"
	"mephi-lab-db/internal/seeder/selector"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type dishRespond struct {
	Results      []models.Dish `json:"results"`
	Offset       uint          `json:"offset"`
	Number       uint          `json:"number"`
	TotalResults uint          `json:"totalResults"`
}

type DishSeeder interface {
	Seed(count uint)
	DummySeed(count uint)
	GetIDs() []uint
	SetCuisineIDs(CuisineIDs []uint)
}

type DishSeederImpl struct {
	db         *gorm.DB
	ids        []uint
	cuisineIDs []uint
}

func NewDishSeeder(db *gorm.DB) DishSeeder {
	return &DishSeederImpl{db: db}
}

func (s *DishSeederImpl) Seed(count uint) {
	keyWords := []string{"fruits", "vegetables", "nuts", "grains", "meat", "fish", "dairy", "oil", "seasoning", "sauce", "breakfast", "lunch", "dinner"}

	var apiCounting uint = 0
	dishes := make([]models.Dish, count)
	//offset := 100
	for i := 0; i < 1; i++ {
		for _, keyWord := range keyWords {
			url := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?apiKey=%s&query=%s&number=%d", apiKey, keyWord, 100)
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

			var responseDish dishRespond
			err = json.Unmarshal(data, &responseDish)
			if err != nil {
				log.Fatal(err)
			}

			for _, currentDish := range responseDish.Results {
				if apiCounting == count {
					break
				}
				dishes[apiCounting] = models.Dish{
					DishID:      apiCounting + 1,
					CuisineID:   selector.NewSelector().RandomSelect(s.cuisineIDs),
					Name:        currentDish.Name,
					EnergyValue: uint(rand.Intn(1200) + 100),
				}
				apiCounting++
			}
		}
	}

	if err := s.db.Create(&dishes).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, dish := range dishes {
		s.ids[i] = dish.DishID
	}
}

func (s *DishSeederImpl) DummySeed(count uint) {
	keyWords := []string{"fruits", "vegetables", "nuts", "grains", "meat", "fish", "dairy", "oil", "seasoning", "sauce", "breakfast", "lunch", "dinner"}

	type DummyDish struct {
		Name string
	}

	type DummyRespond struct {
		Results []DummyDish
	}

	var Dummy DummyRespond
	Dummy.Results = make([]DummyDish, count)
	for i, _ := range Dummy.Results {
		Dummy.Results[i] = DummyDish{
			Name: "Dish " + gofakeit.Noun(),
		}
	}

	var apiCounting uint = 0
	dishes := make([]models.Dish, count)
	//offset := 100
	for i := 0; i < 9; i++ {
		for _, currentDish := range Dummy.Results {
			if apiCounting == count {
				break
			}
			dishes[apiCounting] = models.Dish{
				DishID:      apiCounting + 1,
				CuisineID:   selector.NewSelector().RandomSelect(s.cuisineIDs),
				Name:        currentDish.Name + "with " + keyWords[i],
				EnergyValue: uint(rand.Intn(1200) + 100),
				Time:        uint(rand.Intn(120)),
			}
			apiCounting++
		}
	}

	if err := s.db.Create(&dishes).Error; err != nil {
		panic(err)
	}

	s.ids = make([]uint, count)
	for i, dish := range dishes {
		s.ids[i] = dish.DishID
	}
}

func (s *DishSeederImpl) GetIDs() []uint {
	return s.ids
}

func (s *DishSeederImpl) SetCuisineIDs(CuisineIDs []uint) {
	s.cuisineIDs = CuisineIDs
}
