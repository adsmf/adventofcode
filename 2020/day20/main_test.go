package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 14986175499719
	//Part 2: 2161
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
