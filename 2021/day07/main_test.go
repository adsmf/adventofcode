package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 340987
	//Part 2: 96987874
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestAlternatives(t *testing.T) {
	expectedP1 := 340987
	expectedP2 := 96987874
	type sim = func() (int, int)
	methods := map[string]sim{
		"calcCostsInitial": calcCostsInitial,
		"calcCostsDedup":   calcCostsDedup,
		"calcCostsSlice":   calcCostsSlice,
	}
	for name, fn := range methods {
		t.Run(name, func(t *testing.T) {
			p1, p2 := fn()
			assert.Equal(t, expectedP1, p1)
			assert.Equal(t, expectedP2, p2)
		})
	}
}

func BenchmarkAlternatives(b *testing.B) {
	type sim = func() (int, int)
	methods := map[string]sim{
		"calcCostsInitial": calcCostsInitial,
		"calcCostsDedup":   calcCostsDedup,
		"calcCostsSlice":   calcCostsSlice,
	}
	for name, fn := range methods {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn()
			}
		})
	}
}
