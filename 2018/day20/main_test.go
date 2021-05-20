package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 4214
	//Part 2: 8492
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		regex   string
		longest int
	}
	tests := []testDef{
		{"^WNE$", 3},
		{"^ENWWW(NEEE|SSE(EE|N))$", 10},
		{"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$", 23},
		{"^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$", 31},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			b := newBase(test.regex)
			path := part1(b)
			require.Equal(t, test.longest, path)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
