package main

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 800602729153
	//Part 2: 92173009047076
}

func TestAnswers(t *testing.T) {
}

func TestPart1Examples(t *testing.T) {
	precedence = map[token.Token]int{
		token.ADD: 1,
		token.MUL: 1,
	}
	type testDef struct {
		expression string
		value      int
	}
	tests := []testDef{
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			result2 := calculate([]byte(test.expression))
			assert.Equal(t, test.value, result2)
		})
	}
}

func TestPart2Examples(t *testing.T) {
	precedence = map[token.Token]int{
		token.ADD: 2,
		token.MUL: 1,
	}
	type testDef struct {
		expression string
		value      int
	}
	tests := []testDef{
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			result2 := calculate([]byte(test.expression))
			assert.Equal(t, test.value, result2)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
