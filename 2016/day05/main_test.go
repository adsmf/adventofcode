package main

import "testing"

func ExampleMain() {
	main()
	//Output:
	//Part 1: f77a0e6e
	//Part 2: 999828ec
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
