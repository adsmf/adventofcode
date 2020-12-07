package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Examples(t *testing.T) {
	graph, _ := loadFile("example.txt")
	assert.Equal(t, 32, part2(graph))
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
	_, reverseGraph := loadFile("input.txt")
	for i := 0; i < b.N; i++ {
		part1(reverseGraph)
	}
}

func BenchmarkPart2(b *testing.B) {
	graph, _ := loadFile("input.txt")
	for i := 0; i < b.N; i++ {
		part2(graph)
	}
}
