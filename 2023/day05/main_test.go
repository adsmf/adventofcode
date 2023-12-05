package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 910845529
	//Part 2: 77435348
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
