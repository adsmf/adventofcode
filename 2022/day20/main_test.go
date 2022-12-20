package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 2215
	//Part 2: 8927480683
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
