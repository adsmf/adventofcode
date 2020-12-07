package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Examples(t *testing.T) {
	assert.Equal(t, 32, part2("example.txt"))
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 248
	//Part 2: 57281
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1("input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2("input.txt")
	}
}
