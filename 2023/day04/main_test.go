package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 22897
	//Part 2: 5095824
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
