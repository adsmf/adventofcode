package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Answer(t *testing.T) {
	assert.Equal(t, 15097178, part1())
}

func TestPart2Answer(t *testing.T) {
	assert.Equal(t, 1558663, part2())
}

func ExampleMain() {
	main()
	//Output:
	// Part 1: 15097178
	// Part 2: 1558663
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
