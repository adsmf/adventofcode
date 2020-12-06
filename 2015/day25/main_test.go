package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		col, row int
		expected int
	}
	tests := []testDef{
		{1, 1, 20151125},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			value := part1(20151125, test.col, test.row)
			assert.Equal(t, test.expected, value)
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		// Test structure here
	}
	tests := []testDef{
		// Test data here
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			// Assertions here
		})
	}
}

func TestAnswers(t *testing.T) {
	p1 := part1(20151125, 3029, 2947)
	assert.Equal(t, p1, 19980801)
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 19980801
	//Part 2: -1
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(20151125, 2947, 3029)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2()
	}
}
