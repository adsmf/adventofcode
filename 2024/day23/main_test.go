package main

import (
	"testing"
)

func ExampleMain() {
	main()
	//Output:
	// Part 1: 1238
	// Part 2: bg,bl,ch,fn,fv,gd,jn,kk,lk,pv,rr,tb,vw
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
