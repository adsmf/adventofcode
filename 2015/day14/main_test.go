package main

import (
	"testing"
)

func TestAnswers(t *testing.T) {
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 2660
	//Part 2: 1256
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
