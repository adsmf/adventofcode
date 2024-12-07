package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 12839601725877
	//Part 2: 149956401519484
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
