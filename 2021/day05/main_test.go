package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestAlternatives(t *testing.T) {
	expectedP1 := 6687
	expectedP2 := 19851
	type sim = func() (int, int)
	methods := map[string]sim{
		"loadInputMap":   loadInputMap,
		"loadInputSlice": loadInputSlice,
		"loadInputArray": loadInputArray,
	}
	for name, fn := range methods {
		t.Run(name, func(t *testing.T) {
			p1, p2 := fn()
			assert.Equal(t, expectedP1, p1)
			assert.Equal(t, expectedP2, p2)
		})
	}
}

func BenchmarkGrids(b *testing.B) {
	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loadInputMap()
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
