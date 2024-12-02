package main

import "testing"

func ExampleMain() {
	main()
	//Output:
	//Part 1: 2344935
	//Part 2: 27647262
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkAlternatives(b *testing.B) {
	methods := []struct {
		name string
		fn   func() (int, int)
	}{
		{"Initial", solveInitial},
		{"Mapless", solveMapless},
		{"SingleSlice", solveSingleSlice},
		{"Latest", solve},
	}

	for _, method := range methods {
		b.Run(method.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				method.fn()
			}
		})
	}
}
