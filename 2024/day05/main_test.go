package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 7307
	//Part 2: 4713
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
