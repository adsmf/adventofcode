package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 3065
	//Part 2: 1562536022966
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
