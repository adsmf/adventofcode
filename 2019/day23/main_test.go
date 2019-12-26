package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswers(t *testing.T) {
	assert.Equal(t, 24922, part1())
	assert.Equal(t, 19478, part2())
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 24922
	//Part 2: 19478
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
