package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 480
	//Part 2: 349975
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		pos  int
		dist int
	}
	tests := []testDef{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			dist, _ := getDist(test.pos)
			assert.Equal(t, test.dist, dist)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
