package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Examples(t *testing.T) {

}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 1930, part1())
}

func ExampleMain() {
	main()
	//Output:
	// Part 1: 1930
	// Part 2:
	// ###..####.#..#.#..#.####..##..####.#..#
	// #..#.#....#.#..#..#.#....#..#....#.#..#
	// #..#.###..##...####.###..#......#..#..#
	// ###..#....#.#..#..#.#....#.....#...#..#
	// #....#....#.#..#..#.#....#..#.#....#..#
	// #....#....#..#.#..#.####..##..####..##.
	//
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
