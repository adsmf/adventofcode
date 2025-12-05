package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 617
	//Part 2: 338258295736104
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
