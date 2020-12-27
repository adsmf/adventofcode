package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 1951
	//Part 2: 20002936
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	traps := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countSafe(traps, 40)
	}
}

func BenchmarkPart2(b *testing.B) {
	traps := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countSafe(traps, 400000)
	}
}
