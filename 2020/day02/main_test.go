package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 572
	//Part 2: 306
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
