package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Examples(t *testing.T) {
	type exampleData struct {
		wire1    string
		wire2    string
		distance int
	}
	tests := []exampleData{
		exampleData{
			wire1:    "R8,U5,L5,D3",
			wire2:    "U7,R6,D4,L4",
			distance: 6,
		},
		exampleData{
			wire1:    "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2:    "U62,R66,U55,R34,D71,R55,D58,R83",
			distance: 159,
		},
		exampleData{
			wire1:    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2:    "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			distance: 135,
		},
	}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test %d", id), func(tt *testing.T) {
			g := loadGrid(test.wire1, test.wire2)
			tt.Logf("Grid:\nw1: %#v\nw2: %#v", g.wire1, g.wire2)
			assert.Equal(tt, test.distance, g.findNearestIntersectionDistance())
		})
	}
}

func TestDay2Examples(t *testing.T) {
	type exampleData struct {
		wire1    string
		wire2    string
		distance int
	}
	tests := []exampleData{
		exampleData{
			wire1:    "R8,U5,L5,D3",
			wire2:    "U7,R6,D4,L4",
			distance: 30,
		},
		exampleData{
			wire1:    "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2:    "U62,R66,U55,R34,D71,R55,D58,R83",
			distance: 610,
		},
		exampleData{
			wire1:    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2:    "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			distance: 410,
		},
	}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test %d", id), func(tt *testing.T) {
			g := loadGrid(test.wire1, test.wire2)
			tt.Logf("Grid:\nw1: %#v\nw2: %#v", g.wire1, g.wire2)
			assert.Equal(tt, test.distance, g.findNearestSignalDistance())
		})
	}
}
func TestDay1Answer(t *testing.T) {
	assert.Equal(t, 5319, day1())
}

func TestDay2Answer(t *testing.T) {
	assert.Equal(t, 122514, day2())
}
