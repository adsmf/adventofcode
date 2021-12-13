package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 664
	//Part 2:
	//####.####...##.#..#.####.#....###..#...
	//#....#.......#.#.#.....#.#....#..#.#...
	//###..###.....#.##.....#..#....###..#...
	//#....#.......#.#.#...#...#....#..#.#...
	//#....#....#..#.#.#..#....#....#..#.#...
	//####.#.....##..#..#.####.####.###..####
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestAlternatives(t *testing.T) {
	expectedP1 := 664
	expectedP2 := `####.####...##.#..#.####.#....###..#...
#....#.......#.#.#.....#.#....#..#.#...
###..###.....#.##.....#..#....###..#...
#....#.......#.#.#...#...#....#..#.#...
#....#....#..#.#.#..#....#....#..#.#...
####.#.....##..#..#.####.####.###..####
`
	type folder = func(in string) (int, string)
	methods := map[string]folder{
		"foldIterative":  foldIterative,
		"foldFunctional": foldFunctional,
	}
	for name, fn := range methods {
		t.Run(name, func(t *testing.T) {
			p1, p2 := fn(input)
			assert.Equal(t, expectedP1, p1)
			assert.Equal(t, expectedP2, p2)
		})
	}
}

func BenchmarkAlternatives(b *testing.B) {
	type folder = func(in string) (int, string)
	methods := map[string]folder{
		"foldIterative":  foldIterative,
		"foldFunctional": foldFunctional,
	}
	for name, fn := range methods {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(input)
			}
		})
	}
}
