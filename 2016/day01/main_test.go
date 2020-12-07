package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		input    string
		distance int
	}
	tests := []testDef{
		{"R2, L3", 5},
		{"R2, R2, R2", 2},
		{"R5, L5, R5, R3", 12},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)

			route := load(test.input)
			assert.Equal(t, test.distance, part1(route))
		})
	}
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 246
	//Part 2: 124
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	route := load(string(inputBytes))
	for i := 0; i < b.N; i++ {
		part1(route)
	}
}

func BenchmarkPart2(b *testing.B) {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	route := load(string(inputBytes))
	for i := 0; i < b.N; i++ {
		part2(route)
	}
}
