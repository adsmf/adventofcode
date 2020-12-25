package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 354320
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGenKey(b *testing.B) {
	cPub, dLoop := 12232269, 5882067
	for i := 0; i < b.N; i++ {
		genKey(cPub, dLoop)
	}
}

func BenchmarkGenKeyBig(b *testing.B) {
	cPub, dLoop := 12232269, 5882067
	for i := 0; i < b.N; i++ {
		genKeyBig(cPub, dLoop)
	}
}
