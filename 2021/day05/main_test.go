package main

import (
	"math"
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 6687
	//Part 2: 19851
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGrids(b *testing.B) {
	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loadInput()
		}
	})
	b.Run("slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loadInputSlice()
		}
	})
	b.Run("array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loadInputArray()
		}
	})
}

func BenchmarkOptions(b *testing.B) {
	dX := 109
	dY := 129
	b.Run("math", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = int(math.Max(math.Abs(float64(dX)), math.Abs(float64(dY))))
		}
	})
	b.Run("func", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = max(abs(dX), abs(dY))
		}
	})
}
