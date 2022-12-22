package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 30552
	//Part 2: 184106
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
