package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: ahnofa
	//Part 2: 802
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
