package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 93813115694560
	//Part 2: 3910938071092
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
