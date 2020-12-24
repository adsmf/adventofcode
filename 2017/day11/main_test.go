package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 834
	//Part 2: 1569
}

func TestAnswers(t *testing.T) {
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		route    string
		distance int
	}
	tests := []testDef{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			route := parseRoute(test.route)
			end, _ := traceRoute(point{}, route)
			assert.Equal(t, test.distance, end.movesFrom(point{}))
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
