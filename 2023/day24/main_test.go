package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 12740
	//Part 2: -1
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
