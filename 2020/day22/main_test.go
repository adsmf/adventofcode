package main

import (
	"fmt"
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

func BenchmarkStringKey(b *testing.B) {
	hands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprint(hands)
	}
}

func BenchmarkScore(b *testing.B) {
	hands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scoreHand(hands[0])
		scoreHand(hands[1])
	}
}
