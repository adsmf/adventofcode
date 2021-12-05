package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		inputFile string
		total     int
	}
	tests := []testDef{
		testDef{"examples/part1.txt", 330},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			pairs := loadInput(test.inputFile)
			t.Logf("Map:\n%#v", pairs)
			assert.Equal(t, test.total, pairs.findBest())
		})
	}
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 664
	//Part 2: 640
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2()
	}
}
