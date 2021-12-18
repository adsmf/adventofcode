package main

import (
	"embed"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed examples/*.txt
var examples embed.FS

func ExampleMain() {
	main()
	//Output:
	//Part 1: 4243
	//Part 2: 4701
}

func TestParse(t *testing.T) {
	tests := []string{
		"[[1,2],[[3,4],5]]",
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			parsed := parse(test)
			noSep := strings.ReplaceAll(test, ",", "")
			assert.Equal(t, noSep, parsed.String())
		})
	}
}

func TestExamples(t *testing.T) {
	type testDef struct {
		input     string
		reduced   string
		magnitude int
	}
	tests := []testDef{
		{"[[1,2],[[3,4],5]]", "", 143},
		{"[1,2]\n[3,4]", "[[1,2],[3,4]]", 0},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", "", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", "", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", "", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", "", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", "", 3488},
		{"[[[[4,3],4],4],[7,[[8,4],9]]]\n[1,1]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{getExample("ex4"), "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
		{getExample("ex5"), "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", 4140},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			reduced, mag := part1(test.input)
			if test.reduced != "" {
				noSep := strings.ReplaceAll(test.reduced, ",", "")
				assert.Equal(t, noSep, reduced)
			}
			if test.magnitude > 0 {
				assert.Equal(t, test.magnitude, mag)
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

func getExample(name string) string {
	content, _ := examples.ReadFile("examples/" + name + ".txt")
	return string(content)
}
