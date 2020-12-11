package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	// Part 1: 2281
	// Part 2: 2085
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	floorplan := load("input.txt")
	benchmark = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seatLife(floorplan, false)
	}
}

func BenchmarkPart2(b *testing.B) {
	floorplan := load("input.txt")
	benchmark = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seatLife(floorplan, true)
	}
}
