package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example1.txt
var example1 string

//go:embed example2.txt
var example2 string

func ExampleMain() {
	main()
	//Output:
	//Part 1: 3421
	//Part 2: 84870
}

func TestAnswers(t *testing.T) {
}

func TestExamples(t *testing.T) {
	type testDef struct {
		routeList  string
		expectedP1 int
		expectedP2 int
	}
	tests := []testDef{
		{example1, 10, 36},
		{example2, 19, 103},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			routes := load(test.routeList)
			p1, p2 := explore(routes)
			assert.Equal(t, test.expectedP1, p1)
			assert.Equal(t, test.expectedP2, p2)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
