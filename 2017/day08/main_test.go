package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 5075
	//Part 2: 7310
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
