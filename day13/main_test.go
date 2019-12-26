package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswers(t *testing.T) {
	assert.Equal(t, 341, part1())
	assert.Equal(t, 17138, part2())
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
	// Part 1: 341
	// Part 2: 17138
}
