package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Answer(t *testing.T) {
	assert.Equal(t, 3224742, part1())
}

func TestPart2Answer(t *testing.T) {
	assert.Equal(t, 7960, part2())
}

func ExampleMain() {
	main()
	//Output:
	// Part 1: 3224742
	// Part 2: 7960
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
