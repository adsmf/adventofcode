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
	gameState := gameState{
		playerHand{1, 2, 3, 4, 5, 6, 7, 8, 9},
		playerHand{21, 22, 23, 24, 25, 26, 27, 28, 29},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprint(gameState)
	}
}

func BenchmarkScore(b *testing.B) {
	gameState := gameState{
		playerHand{1, 2, 3, 4, 5, 6, 7, 8, 9},
		playerHand{21, 22, 23, 24, 25, 26, 27, 28, 29},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scoreHand(gameState[0])
		scoreHand(gameState[1])
	}
}
