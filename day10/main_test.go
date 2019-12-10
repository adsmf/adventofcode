package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type example struct {
		inputFile    string
		numAsteroids int
		bestLocation vector
		bestCount    int
	}
	examples := []example{
		example{
			inputFile:    "part1ex1.txt",
			numAsteroids: 10,
			bestLocation: vector{3, 4},
			bestCount:    8,
		},
	}
	for id, test := range examples {
		t.Run(fmt.Sprintf("Day 10 part 1 - %d", id), func(t *testing.T) {
			g := loadInputFile("examples/" + test.inputFile)
			t.Logf("%d %v", len(g.asteroids), g)
			ast, count := g.getBestLineOfSight()
			assert.EqualValues(t, test.bestLocation, ast.position)
			assert.Equal(t, test.bestCount, count)
		})
	}
}

func TestPart2Examples(t *testing.T) {
	base := asteroid{vector{8, 3}}
	removalOrder := []asteroid{
		asteroid{vector{8, 1}},
		asteroid{vector{9, 0}},
		asteroid{vector{9, 1}},
		asteroid{vector{10, 0}},
		asteroid{vector{9, 2}},
	}
	g := loadInputFile("examples/part2ex1.txt")
	destroyed := destroyN(g, base, len(removalOrder))
	for num, expectDestroyed := range removalOrder {
		t.Run(fmt.Sprintf("Test destroy order %d", num+1), func(t *testing.T) {
			assert.Equal(t, expectDestroyed, destroyed[num])
		})
	}
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 221
	//Part 2: 806
}
