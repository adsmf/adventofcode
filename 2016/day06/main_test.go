package main

import "testing"

func ExampleMain() {
	main()
	//Output:
	//Part 1: qqqluigu
	//Part 2: lsoypmia
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
