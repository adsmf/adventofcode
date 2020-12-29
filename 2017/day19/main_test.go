package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: QPRYCIOLU
	//Part 2: 16162
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
