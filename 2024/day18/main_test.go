package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 454
	//Part 2: 8,51
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
