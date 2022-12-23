package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 4249
	//Part 2: 980
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestPoints(t *testing.T) {
	type testData struct{ x, y int }
	tests := []testData{
		{0, 0},
		{1, 1},
		{-1, -1},
		{(1<<(axisBits-1) - 18), (1<<(axisBits-1) - 27)},
		{(1<<(axisBits-1) - 18) * -1, (1<<(axisBits-1) - 27) * -1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test), func(t *testing.T) {
			pos := pointAt(test.x, test.y)
			result := testData{pos.x(), pos.y()}
			assert.Equal(t, test, result)
		})
	}
}
