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
	//Part 2: -1
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
