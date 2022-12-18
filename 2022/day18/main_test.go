package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 4302
	//Part 2: 2492
}

func TestPoint(t *testing.T) {
	type testPoint [3]int
	testPoints := []testPoint{
		{1, 0, 0}, {-1, 0, 0},
		{0, 1, 0}, {0, -1, 0},
		{0, 0, 1}, {0, 0, -1},
	}
	for _, test := range testPoints {
		t.Run(fmt.Sprint(test), func(t *testing.T) {
			pos := pointAt(test[0], test[1], test[2])
			assert.Equal(t, test[0], pos.x())
			assert.Equal(t, test[1], pos.y())
			assert.Equal(t, test[2], pos.z())
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
