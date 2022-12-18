package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 4302
	//Part 2: 2492
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
