package j

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedDict(t *testing.T) {
	t.Log(`
	Given a JSON string representing a list of products with their categories, 
	write a function to parse the JSON and create a nested dictionary representing the category hierarchy.
	The function should return the nested dictionary where each category is a key, 
	and the corresponding value is either another nested dictionary representing its subcategories
	or an empty dictionary if it has no subcategories. 
	Each product has a "name" field representing its name, 
	and a "category" field representing its category.
`)

	jsonStr := `
[
  {
    "name": "Product 1",
    "category": "Electronics"
  },
  {
    "name": "Product 2",
    "category": "Electronics/Computers"
  },
  {
    "name": "Product 3",
    "category": "Electronics/Phones"
  },
  {
    "name": "Product 4",
    "category": "Clothing"
  },
  {
    "name": "Product 5",
    "category": "Clothing/Men"
  },
  {
    "name": "Product 6",
    "category": "Clothing/Women"
  }
]
`
	exp := map[string]interface{}{
		"Clothing": map[string]interface{}{
			"Men":   map[string]interface{}{},
			"Women": map[string]interface{}{},
		},
		"Electronics": map[string]interface{}{
			"Computers": map[string]interface{}{},
			"Phones":    map[string]interface{}{},
		},
	}
	assert.Equal(t, exp, RunNestedDict(jsonStr))
}

type NameCategory []struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func RunNestedDict(jsonStr string) map[string]interface{} {
	var in NameCategory
	err := json.Unmarshal([]byte(jsonStr), &in)
	if err != nil {
		panic(err)
	}

	categoryHierarchy := make(map[string]interface{})
	for _, v := range in {
		categoryPath := strings.Split(v.Category, "/")
		currentDict := categoryHierarchy

		for _, category := range categoryPath {
			if _, ok := currentDict[category]; !ok {
				currentDict[category] = make(map[string]interface{})
			}
			currentDict = currentDict[category].(map[string]interface{})
		}
	}
	return categoryHierarchy
}
