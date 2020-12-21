package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 2635
	//Part 2: xncgqbcp,frkmp,qhqs,qnhjhn,dhsnxr,rzrktx,ntflq,lgnhmx
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
