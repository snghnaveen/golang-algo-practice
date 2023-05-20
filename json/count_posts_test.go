package j

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPosts(t *testing.T) {
	t.Log(`
	Given a JSON string representing a list of blog posts and their comments, 
	write a function to parse the JSON and generate a report that includes the total number of 
	comments for each blog post. The JSON string contains an array of objects, 
	where each object represents a blog post and has a "title" field representing 
	the title of the post, and a "comments" field containing a list of comments.
	`)

	jsonStr := `
[
  {
    "title": "Post 1",
    "comments": [
      {
        "user": "User 1",
        "comment": "Comment 1"
      },
      {
        "user": "User 2",
        "comment": "Comment 2"
      },
      {
        "user": "User 3",
        "comment": "Comment 3"
      }
    ]
  },
  {
    "title": "Post 2",
    "comments": [
      {
        "user": "User 4",
        "comment": "Comment 4"
      },
      {
        "user": "User 5",
        "comment": "Comment 5"
      }
    ]
  },
  {
    "title": "Post 3",
    "comments": [
      {
        "user": "User 6",
        "comment": "Comment 6"
      },
      {
        "user": "User 7",
        "comment": "Comment 7"
      },
      {
        "user": "User 8",
        "comment": "Comment 8"
      },
      {
        "user": "User 9",
        "comment": "Comment 9"
      }
    ]
  }
]
`
	exp := map[string]int{
		"Post 1": 3,
		"Post 2": 2,
		"Post 3": 4,
	}

	assert.Equal(t, exp, RunCountPosts(jsonStr))

}

func RunCountPosts(jsonStr string) map[string]int {

	type Comments struct {
		User    string `json:"user"`
		Comment string `json:"comment"`
	}
	type Posts []struct {
		Title    string     `json:"title"`
		Comments []Comments `json:"comments"`
	}

	var posts Posts
	err := json.Unmarshal([]byte(jsonStr), &posts)
	if err != nil {
		panic(err)
	}

	out := make(map[string]int)
	for _, v := range posts {
		out[v.Title] = len(v.Comments)
	}

	return out
}
