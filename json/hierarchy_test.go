package j

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHierarchy(t *testing.T) {
	t.Log(`
	Given a JSON string representing a hierarchical organizational structure, 
	write a function to parse the JSON and print the hierarchy in a tree-like format. 
	Each employee has a "name" field, and there is a "reports" field containing a list of their direct reports.
	`)

	jsonStr := `
		{
			"name": "John Doe",
			"reports": [
			{
				"name": "Jane Smith",
				"reports": [
				{
					"name": "Tom Johnson",
					"reports": []
				},
				{
					"name": "Emily Davis",
					"reports": []
				}
				]
			},
			{
				"name": "Mike Brown",
				"reports": [
				{
					"name": "Sarah Wilson",
					"reports": []
				},
				{
					"name": "Mark Lee",
					"reports": [
					{
						"name": "Alex Clark",
						"reports": []
					}
					]
				}
				]
			}
			]
		}
	`

	exp := `
- John Doe
  - Jane Smith
    - Tom Johnson
    - Emily Davis
  - Mike Brown
    - Sarah Wilson
    - Mark Lee
      - Alex Clark
`

	assert.Equal(t, exp, RunHierarchy(jsonStr).String())
}

type Reports struct {
	Name    string    `json:"name"`
	Reports []Reports `json:"reports"`
}

func RunHierarchy(jsonStr string) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	var in Reports
	err := json.Unmarshal([]byte(jsonStr), &in)
	if err != nil {
		panic(err)
	}

	buffer.WriteString("\n")
	PrintHierarchy(buffer, in, "- ")
	return buffer
}

func PrintHierarchy(b *bytes.Buffer, in Reports, indent string) *bytes.Buffer {
	b.WriteString(indent + in.Name + "\n")
	indent = "  " + indent
	for _, v := range in.Reports {
		PrintHierarchy(b, v, indent)
	}
	return b
}
