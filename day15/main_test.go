package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Answer(t *testing.T) {
	assert.Equal(t, 252, part1())
}

func TestPart2Answer(t *testing.T) {
	assert.Equal(t, 350, part2())
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

func ExampleMain() {
	main()
	//Output:
	// Part 1: 252
	// Part 2: 350
}
