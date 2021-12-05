package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 11754
	//Part 2: 4098
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
