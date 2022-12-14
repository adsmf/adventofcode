package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 768
	//Part 2: 26686
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
