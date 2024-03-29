package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 1427868
	//Part 2: 1568138742
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkMethods(b *testing.B) {
	b.Run("Initial", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			followRouteInitial()
		}
	})
	b.Run("Fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			followRouteFast()
		}
	})
}
