package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 248113761
	//Part 2: 246285222
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		hand   handInfo
		expect handType
	}
	tests := []testDef{
		{handInfo{cards: "32T3K"}, handOnePair},
		{handInfo{cards: "KK677"}, handTwoPair},
		{handInfo{cards: "KTJJT"}, handTwoPair},
		{handInfo{cards: "T55J5"}, handThreeOfKind},
		{handInfo{cards: "QQQJA"}, handThreeOfKind},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			assert.Equal(t, test.expect, test.hand.score(false))
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		hand   handInfo
		expect handType
	}
	tests := []testDef{
		{handInfo{cards: "32T3K"}, handOnePair},
		{handInfo{cards: "KK677"}, handTwoPair},
		{handInfo{cards: "T55J5"}, handFourOfKind},
		{handInfo{cards: "KTJJT"}, handFourOfKind},
		{handInfo{cards: "QQQJA"}, handFourOfKind},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part2-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			assert.Equal(t, test.expect, test.hand.score(true))
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
