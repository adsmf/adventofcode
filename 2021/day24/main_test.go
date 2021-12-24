package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 99691891979938
	//Part 2: 27141191213911
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
