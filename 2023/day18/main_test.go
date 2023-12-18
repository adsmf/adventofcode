package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 50746
	//Part 2: 70086216556038
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
