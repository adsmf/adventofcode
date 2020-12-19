package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 326618
	//Part 2: 21841249
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
