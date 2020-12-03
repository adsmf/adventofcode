package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 262
	//Part 2: 2698900776
}

func ExampleMainAlt() {
	mainAlt()
	//Output:
	//Part 1: 262
	//Part 2: 2698900776
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

func BenchmarkPart1alt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1alt()
	}
}

func BenchmarkPart2alt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2alt()
	}
}
