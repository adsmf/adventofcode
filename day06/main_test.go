package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type test struct {
		definitions []string
		orbits      int
	}
	tests := []test{
		test{
			definitions: []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
			orbits:      42,
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part 1 example %d", id), func(t *testing.T) {
			g := newGraph(test.definitions)
			t.Logf("Graph:\n%#v", g)
			assert.Equal(t, test.orbits, g.countOrbits())
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type test struct {
		definitions []string
		start, end  string
		length      int
	}
	tests := []test{
		test{
			definitions: []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"},
			start:       "YOU",
			end:         "SAN",
			length:      4,
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part 1 example %d", id), func(t *testing.T) {
			g := newGraph(test.definitions)
			t.Logf("Graph:\n%#v", g)
			assert.Equal(t, test.length, g.calcRouteLength(test.start, test.end))
		})
	}
}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 295936, part1())
	assert.Equal(t, 457, part2())
}

func TestMainRuns(t *testing.T) {
	assert.NotPanics(t, func() { main() })
}
