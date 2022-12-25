package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 20-==01-2-=1-2---1-0
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
