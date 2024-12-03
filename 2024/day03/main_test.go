package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 185797128
	//Part 2: 89798695
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
