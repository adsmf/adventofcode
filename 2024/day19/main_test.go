package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 340
	//Part 2: 717561822679428
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
