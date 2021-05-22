package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 120
	//Part 2: 2204099
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
