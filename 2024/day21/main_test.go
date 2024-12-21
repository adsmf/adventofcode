package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 219366
	//Part 2: 271631192020464
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
