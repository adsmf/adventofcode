package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 17765746710228
	//Part 2: 4401465949086
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	commands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(commands)
	}
}

func BenchmarkPart2(b *testing.B) {
	commands := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(commands)
	}
}
