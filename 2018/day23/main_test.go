package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 889
	//Part 2: 160646364
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
