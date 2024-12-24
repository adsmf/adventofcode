package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 55920211035878
	//Part 2: btb,cmv,mwp,rdg,rmj,z17,z23,z30
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
