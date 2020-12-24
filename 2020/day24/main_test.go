package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 312
	//Part 2: 3733
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		layTiles("input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	_, tiles := layTiles("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tileLife(tiles, 100)
	}
}
