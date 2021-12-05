package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 6687
	//Part 2: 19851
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
