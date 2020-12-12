package main

import (
	"testing"

	"github.com/adsmf/adventofcode/utils"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 1441
	//Part 2: 61616
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.ReadInputLines("input.txt")
	}
}

func BenchmarkPart1(b *testing.B) {
	commands := utils.ReadInputLines("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(commands)
	}
}

func BenchmarkPart2(b *testing.B) {
	commands := utils.ReadInputLines("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(commands)
	}
}
