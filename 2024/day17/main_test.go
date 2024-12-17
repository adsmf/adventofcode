package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 1,4,6,1,6,4,3,0,3
	//Part 2: 265061364597659
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
