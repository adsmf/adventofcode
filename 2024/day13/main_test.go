package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 32041
	//Part 2: 95843948914827
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
