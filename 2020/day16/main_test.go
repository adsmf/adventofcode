package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 25895
	//Part 2: 5865723727753
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
