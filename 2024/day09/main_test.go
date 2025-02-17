package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 6421128769094
	//Part 2: 6448168620520
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
