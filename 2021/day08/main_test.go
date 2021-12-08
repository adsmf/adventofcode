package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 264
	//Part 2: 1063760
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
