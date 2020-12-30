package main

import (
	"testing"
)

func ExampleMain() {
	isTest = true
	main()
	//Output:
	//Part 1: 31471
	//Part 2: 24169
}

func BenchmarkMain(b *testing.B) {
	benchmark, isTest = true, true
	for i := 0; i < b.N; i++ {
		main()
	}
}
