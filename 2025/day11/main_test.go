package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 555
	//Part 2: 502447498690860
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
