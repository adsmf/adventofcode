package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 386
	//Part 2: 2276
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	initial := loadInitial("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(initial)
	}
}

func BenchmarkPart2(b *testing.B) {
	initial := loadInitial("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(initial)
	}
}
