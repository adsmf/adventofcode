package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	// Part 1: 1759
	// Part 2: 1805
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkFindIncreases(b *testing.B) {
	depths := getDepths()
	b.Run("part1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			findIncreases(1, depths)
		}
	})
	b.Run("part2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			findIncreases(1, depths)
		}
	})
}

func BenchmarkGetDepths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDepths()
	}
}
