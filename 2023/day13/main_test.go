package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 34911
	//Part 2: 33183
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
