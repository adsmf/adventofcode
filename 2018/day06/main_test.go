package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 2342
	//Part 2: 43302
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
