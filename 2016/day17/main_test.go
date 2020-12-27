package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: RDDRLDRURD
	//Part 2: 448
}

func TestAnswers(t *testing.T) {
}

func TestExamples(t *testing.T) {
	type testDef struct {
		input  string
		path   string
		length int
	}
	tests := []testDef{
		{"ihgpwlah", "DDRRRD", 370},
		{"kglvqrro", "DDUDRLRRUDRD", 492},
		{"ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR", 830},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			p1, p2 := solve(test.input)
			assert.Equal(t, test.path, p1)
			assert.Equal(t, test.length, p2)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
