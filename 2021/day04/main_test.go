package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 39984
	//Part 2: 8468
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkMethods(b *testing.B) {
	b.Run("LoadInput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loadInput()
		}
	})
	draws, boards := loadInput()
	b.Run("GenerateWinSets", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateWinSets(draws, boards)
		}
	})
	b.Run("getScores", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			getScores(draws, boards)
		}
	})
}
