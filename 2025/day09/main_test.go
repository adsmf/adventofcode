package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 4758121828
	//Part 2: 1577956170
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
