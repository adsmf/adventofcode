package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	state := setup()
	expectedScores := []recipe{3, 7}

	for idx, score := range state.board {
		assert.Equal(t, expectedScores[idx], score)
	}

	expectedCurRecipe := []elf{0, 1}
	for idx, curRecipe := range state.elves {
		assert.Equal(t, expectedCurRecipe[idx], curRecipe)
	}
}

func TestTickOnce(t *testing.T) {
	expectedScores := scoreboard{3, 7, 1, 0}
	expectedElves := workers{0, 1}

	state := setup()
	state = tick(state)

	assert.Equal(t, expectedScores.AsString(), state.board.AsString())
	assert.Equal(t, expectedElves.AsString(), state.elves.AsString())
}

func TestP1Example(t *testing.T) {
	logger = t.Logf
	states := []string{
		"(3)[7]",
		"(3)[7] 1  0 ",
		" 3  7  1 [0](1) 0 ",
		" 3  7  1  0 [1] 0 (1)",
		"(3) 7  1  0  1  0 [1] 2 ",
		" 3  7  1  0 (1) 0  1  2 [4]",
		" 3  7  1 [0] 1  0 (1) 2  4  5 ",
		" 3  7  1  0 [1] 0  1  2 (4) 5  1 ",
		" 3 (7) 1  0  1  0 [1] 2  4  5  1  5 ",
		" 3  7  1  0  1  0  1  2 [4](5) 1  5  8 ",
		" 3 (7) 1  0  1  0  1  2  4  5  1  5  8 [9]",
		" 3  7  1  0  1  0  1 [2] 4 (5) 1  5  8  9  1  6 ",
		" 3  7  1  0  1  0  1  2  4  5 [1] 5  8  9  1 (6) 7 ",
		" 3  7  1  0 (1) 0  1  2  4  5  1  5 [8] 9  1  6  7  7 ",
		" 3  7 [1] 0  1  0 (1) 2  4  5  1  5  8  9  1  6  7  7  9 ",
		" 3  7  1  0 [1] 0  1  2 (4) 5  1  5  8  9  1  6  7  7  9  2 ",
	}

	var state cookState
	for tickID, expectedState := range states {
		t.Run(fmt.Sprintf("Tick %d", tickID), func(t *testing.T) {
			if tickID == 0 {
				state = setup()
			} else {
				state = tick(state)
			}
			assert.Equal(t, expectedState, state.AsString())
		})
	}
}

func TestExampleNextTen(t *testing.T) {
	examples := map[int]string{
		9:    "5158916779",
		18:   "9251071085",
		2018: "5941429882",
	}
	for startAt, expected := range examples {
		state := setup()
		result := nextTen(state, startAt)
		assert.Equal(t, expected, result)
	}
}

func TestExampleFirstAppearances(t *testing.T) {
	examples := map[string]int{
		"51589": 9,
		"01245": 5,
		"92510": 18,
		"59414": 2018,
	}
	for target, expected := range examples {
		t.Run(target, func(t *testing.T) {
			state := setup()
			result := firstAppears(state, target)
			assert.Equal(t, expected, result)
		})
	}
}
