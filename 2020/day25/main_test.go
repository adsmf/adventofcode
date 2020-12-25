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

func BenchmarkFindLoop(b *testing.B) {
	doorPublic := 19452773
	for i := 0; i < b.N; i++ {
		findLoop(doorPublic)
	}
}

func BenchmarkFindLoopBSGS(b *testing.B) {
	doorPublic := 19452773
	for i := 0; i < b.N; i++ {
		findLoopBSGS(doorPublic)
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
