package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 386640
	//Part 2: 1733403626279
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
