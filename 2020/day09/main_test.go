package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	// Part 1: 1492208709
	// Part 2: 238243506
	// Part 2 alt: 238243506
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	p1 := part1()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(p1)
	}
}

func BenchmarkPart2alt(b *testing.B) {
	p1 := part1()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2alt(p1)
	}
}
