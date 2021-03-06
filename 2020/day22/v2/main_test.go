package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 33434
	//Part 2: 31657
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
	for i := 0; i < b.N; i++ {
		part2()
	}
}

func BenchmarkHashSprint(b *testing.B) {
	hands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hands.hashSprint()
	}
}

func BenchmarkHashScore(b *testing.B) {
	hands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hands.hashScore()
	}
}

func BenchmarkHashByteSlice(b *testing.B) {
	hands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hands.hashFromByteSlice()
	}
}

func BenchmarkHashBuffer(b *testing.B) {
	hands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hands.hashByteBuffer()
	}
}

func BenchmarkHashCombine(b *testing.B) {
	hands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hands.hashCombine()
	}
}
