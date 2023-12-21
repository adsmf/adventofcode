package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 3764
	//Part 2: 622926941971282
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
