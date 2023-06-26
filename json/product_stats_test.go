package j

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductStats(t *testing.T) {
	t.Log(`
	Given a JSON string representing a collection of products and their attributes, 
	write a function to parse the JSON and generate a report that includes the average 
	price and average rating for each category of products. 
	The JSON string contains an array of objects, 
	where each object represents a product and has 
	a "name" field representing the product name, 
	a "category" field representing the product category, 
	a "price" field representing the product price, and 
	a "rating" field representing the product rating.
	`)

	jsonStr := `
[
  {
    "name": "Product 1",
    "category": "Electronics",
    "price": 1000,
    "rating": 4.5
  },
  {
    "name": "Product 2",
    "category": "Electronics",
    "price": 1500,
    "rating": 4.2
  },
  {
    "name": "Product 3",
    "category": "Clothing",
    "price": 50,
    "rating": 3.9
  },
  {
    "name": "Product 4",
    "category": "Clothing",
    "price": 80,
    "rating": 4.1
  },
  {
    "name": "Product 5",
    "category": "Electronics",
    "price": 1200,
    "rating": 4.7
  },
  {
    "name": "Product 6",
    "category": "Home",
    "price": 200,
    "rating": 4.0
  },
  {
    "name": "Product 7",
    "category": "Home",
    "price": 300,
    "rating": 4.3
  }
]	
`
	exp := map[string]map[string]float64{
		"Electronics": {
			"avg_price":  1233.33,
			"avg_rating": 4.47,
		},
		"Clothing": {
			"avg_price":  65.00,
			"avg_rating": 4.00,
		},
		"Home": {
			"avg_price":  250.00,
			"avg_rating": 4.15,
		},
	}

	assert.Equal(t, exp, RunProductStats(jsonStr))

}

func RunProductStats(jsonStr string) map[string]map[string]float64 {
	out := make(map[string]map[string]float64)

	type Product struct {
		Name     string  `json:"name"`
		Category string  `json:"category"`
		Price    float64 `json:"price"`
		Rating   float64 `json:"rating"`
	}

	var products []Product
	err := json.Unmarshal([]byte(jsonStr), &products)
	if err != nil {
		panic(err)
	}

	categoryStats := make(map[string]struct {
		Count       int
		TotalPrice  float64
		TotalRating float64
	})

	for _, p := range products {
		category := p.Category
		categoryStat := categoryStats[category]

		categoryStat.Count++
		categoryStat.TotalPrice = categoryStat.TotalPrice + p.Price
		categoryStat.TotalRating = categoryStat.TotalRating + p.Rating

		categoryStats[category] = categoryStat
	}

	for category, stats := range categoryStats {
		count := stats.Count
		avgPrice := stats.TotalPrice / float64(count)
		avgRating := stats.TotalRating / float64(count)

		out[category] = map[string]float64{
			"avg_price":  math.Round(avgPrice*100) / 100,
			"avg_rating": math.Round(avgRating*100) / 100,
		}
	}

	return out
}
