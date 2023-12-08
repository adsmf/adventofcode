package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 19241
	//Part 2: 9606140307013
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
