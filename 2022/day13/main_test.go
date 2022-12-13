package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 5717
	//Part 2: 25935
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		list1, list2 string
		expect       bool
	}
	tests := []testDef{
		{"[1,1,3,1,1]", "[1,1,5,1,1]", true},
		{"[[1],[2,3,4]]", "[[1],4]", true},
		{"[9]", "[[8,7,6]]", false},
		{"[[4,4],4,4]", "[[4,4],4,4,4]", true},
		{"[7,7,7,7]", "[7,7,7]", false},
		{"[]", "[3]", true},
		{"[[[]]]", "[[]]", false},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]", false},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			cmp := compare([]byte(test.list1), []byte(test.list2))
			assert.Equal(t, test.expect, cmp)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
