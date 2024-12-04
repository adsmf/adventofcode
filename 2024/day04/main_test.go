package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 2464
	//Part 2: 1982
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
