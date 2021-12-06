package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 386640
	//Part 2: 1733403626279
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
