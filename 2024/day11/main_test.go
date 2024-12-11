package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 198075
	//Part 2: 235571309320764
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input  int
		even   bool
		e1, e2 int
	}{
		{1, false, 0, 0},
		{10, true, 1, 0},
		{11, true, 1, 1},
		{110, false, 0, 0},
		{253000, true, 253, 0},
	}
	for _, test := range tests {
		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			a1, a2, even := split(test.input)
			require.Equal(t, test.even, even)
			if test.even {
				require.Equal(t, test.e1, a1)
				require.Equal(t, test.e2, a2)
			}
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
