package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		playerHP      int
		playerMana    int
		bossHP        int
		bossDamage    int
		expectedSpend int
	}
	tests := []testDef{
		{10, 250, 13, 8, 226},
		{10, 250, 14, 8, 641},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			game := gameData{
				playerNext: true,
				boss: entity{
					hitpoints:  test.bossHP,
					baseDamage: test.bossDamage,
				},
				player: entity{
					hitpoints: test.playerHP,
					mana:      test.playerMana,
				},
			}
			mana := game.optimana()
			assert.Equal(t, test.expectedSpend, mana)
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		// Test structure here
	}
	tests := []testDef{
		// Test data here
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			// Assertions here
		})
	}
}

func TestAnswers(t *testing.T) {
	debug = false
	p1 := part1()
	assert.Greater(t, p1, 742)
	assert.Greater(t, p1, 777)
	assert.Greater(t, p1, 844)
	assert.NotEqual(t, 1309, p1)
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 1269
	//Part 2: 1309
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2()
	}
}
