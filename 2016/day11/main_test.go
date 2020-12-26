package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 31
	//Part 2: 55
}

func TestHash(t *testing.T) {
	states := map[facilityHash]facilityState{}
	for elevator := 1; elevator <= 4; elevator++ {
		for e1r := 1; e1r <= 4; e1r++ {
			for e1c := 1; e1c <= 4; e1c++ {
				for e2r := 1; e2r <= 4; e2r++ {
					for e2c := 1; e2c <= 4; e2c++ {
						state := facilityState{
							elevatorFloor: elevator,
							rtgs:          []int{e1r, e2r},
							chips:         []int{e1c, e2c},
						}
						hash := state.hash()
						if prev, found := states[hash]; found {
							assert.Fail(t, fmt.Sprintf("Collision:\nHash: %d\nCur: %v\nPrev: %v", hash, state, prev))
						}
						states[hash] = state
					}
				}
			}
		}
	}
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

func BenchmarkHash(b *testing.B) {
	state := load("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		state.hash()
	}
}
