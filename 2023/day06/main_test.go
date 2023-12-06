package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 114400
	//Part 2: 21039729
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
