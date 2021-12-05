package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 143415
	//Part 2: 49944
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
