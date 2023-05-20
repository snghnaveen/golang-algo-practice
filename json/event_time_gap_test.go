package j

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventTimeGap(t *testing.T) {
	t.Log(`
	Given a JSON string representing a collection of events and their timestamps, 
	write a function to parse the JSON and generate a report that includes 
	the maximum time gap between consecutive events for each event type. 
	The JSON string contains an array of objects, 
	where each object represents an event and 
	has a "type" field representing the event type, 
	and a "timestamp" field representing the event timestamp in milliseconds.
	`)

	jsonStr := `
[
  {
    "type": "Event 1",
    "timestamp": 1621479600000
  },
  {
    "type": "Event 2",
    "timestamp": 1621479700000
  },
  {
    "type": "Event 1",
    "timestamp": 1621479800000
  },
  {
    "type": "Event 3",
    "timestamp": 1621479850000
  },
  {
    "type": "Event 2",
    "timestamp": 1621479900000
  },
  {
    "type": "Event 1",
    "timestamp": 1621479950000
  }
]
`

	exp := map[string]int64{
		"Event 1": 100000,
		"Event 2": 100000,
		"Event 3": 50000,
	}

	assert.Equal(t, exp, RunEventTimeGap(jsonStr))
}

func RunEventTimeGap(jsonStr string) map[string]int64 {
	type EventTime []struct {
		Type      string `json:"type"`
		Timestamp int64  `json:"timestamp"`
	}

	var eventTime EventTime

	err := json.Unmarshal([]byte(jsonStr), &eventTime)
	if err != nil {
		panic(err)
	}

	maxTimeGaps := make(map[string]int64)

	// Sort events by timestamp
	sort.Slice(eventTime, func(i, j int) bool {
		return eventTime[i].Timestamp < eventTime[j].Timestamp
	})

	for i := 1; i < len(eventTime); i++ {
		currentEvent := eventTime[i]
		previousEvent := eventTime[i-1]

		timeGap := currentEvent.Timestamp - previousEvent.Timestamp

		if timeGap > maxTimeGaps[currentEvent.Type] {
			maxTimeGaps[currentEvent.Type] = timeGap
		}
	}

	return maxTimeGaps
}
