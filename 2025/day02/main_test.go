package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 29818212493
	//Part 2: 37432260594
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
