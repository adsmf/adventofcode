package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 817896682
	//Part 2: 250924073918341
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
