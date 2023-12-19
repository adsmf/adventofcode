package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 489392
	//Part 2: 134370637448305
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
