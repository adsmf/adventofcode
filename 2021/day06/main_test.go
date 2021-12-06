package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 386640
	//Part 2: 1733403626279
}

func TestAlternatives(t *testing.T) {
	expectedP1 := 386640
	expectedP2 := 1733403626279
	type sim = func() (int, int)
	methods := map[string]sim{
		"runSim":              runSim,
		"runSimRing":          runSimRing,
		"runSimNoAppend":      runSimNoAppend,
		"runSimNoPreallocate": runSimNoPreallocate,
		"runSimSlice":         runSimSlice,
	}
	for name, fn := range methods {
		t.Run(name, func(t *testing.T) {
			p1, p2 := fn()
			assert.Equal(t, expectedP1, p1)
			assert.Equal(t, expectedP2, p2)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkRing(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		runSimRing()
	}
}

func BenchmarkSims(b *testing.B) {
	b.Run("circular", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runSim()
		}
	})
	b.Run("slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runSimSlice()
		}
	})
	b.Run("noPreallocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runSimNoPreallocate()
		}
	})
	b.Run("noAppend", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runSimNoAppend()
		}
	})
	b.Run("ring", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runSimRing()
		}
	})
}
