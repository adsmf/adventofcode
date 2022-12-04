package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 477
	//Part 2: 830
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
