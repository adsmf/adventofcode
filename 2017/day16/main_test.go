package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: iabmedjhclofgknp
	//Part 2: oildcmfeajhbpngk
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	ops := choreograph("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dance(ops, 1)
	}
}

func BenchmarkPart2(b *testing.B) {
	ops := choreograph("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dance(ops, 1000000000)
	}
}
