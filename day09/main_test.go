package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type game struct {
	players       int
	marbles       int
	expectedScore int
}

func TestExampleGames(t *testing.T) {
	// Examples given:
	//   10 players; last marble is worth 1618 points: high score is 8317
	//   13 players; last marble is worth 7999 points: high score is 146373
	//   17 players; last marble is worth 1104 points: high score is 2764
	//   21 players; last marble is worth 6111 points: high score is 54718
	//   30 players; last marble is worth 5807 points: high score is 37305
	examples := []game{
		game{10, 1618, 8317},
		game{13, 7999, 146373},
		game{17, 1104, 2764},
		game{21, 6111, 54718},
		game{30, 5807, 37305},
	}
	for exampleID, example := range examples {
		t.Run(fmt.Sprintf("Example %d", exampleID), func(t *testing.T) {
			t.Logf("Running example %+v", example)
			score := playMarbles(example.players, example.marbles)
			assert.Equal(t, example.expectedScore, score)
		})
	}
}
