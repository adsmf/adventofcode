package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5Part1Examples(t *testing.T) {
	debug = noOut
	tests := map[string]string{
		// Previous
		"1,0,0,0,99":                    "2,0,0,0,99",
		"2,3,0,3,99":                    "2,3,0,6,99",
		"2,4,4,5,99,0":                  "2,4,4,5,99,9801",
		"1,1,1,4,99,5,6,0,99":           "30,1,1,4,2,5,6,0,99",
		"1,9,10,3,2,3,11,0,99,30,40,50": "3500,9,10,70,2,3,11,0,99,30,40,50",
		// New
		"1002,4,3,4,33": "1002,4,3,4,99",
		// Mine
		"3,1,99": "3,1002,99",
	}

	for program, expected := range tests {
		inputs := make(chan int, 1)
		outputs := make(chan int)
		var output int
		go func() {
			output = <-outputs
			t.Logf("Got output: %d", output)
		}()
		inputs <- 1002
		mach := newMachine(program, inputs, outputs)
		assert.Equal(t, 0, output)
		mach.run()
		assert.Equal(t, expected, mach.String())
	}
}

func TestDay5Part2Examples(t *testing.T) {
	type part2inputtest map[int]int

	tests := map[string]part2inputtest{
		"3,9,8,9,10,9,4,9,99,-1,8": map[int]int{
			7: 0,
			8: 1,
			9: 0,
		},
		"3,9,7,9,10,9,4,9,99,-1,8": map[int]int{
			7: 1,
			8: 0,
			9: 0,
		},
		"3,3,1108,-1,8,3,4,3,99": map[int]int{
			7: 0,
			8: 1,
			9: 0,
		},
		"3,3,1107,-1,8,3,4,3,99": map[int]int{
			7: 1,
			8: 0,
			9: 0,
		},
		//  1 2  3  4 5  6  7  8 9 10 11 12131415
		"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9": map[int]int{
			//                           xx
			-1: 1,
			0:  0,
			1:  1,
			2:  1,
		},
		"3,3,1105,-1,9,1101,0,0,12,4,12,99,1": map[int]int{
			-1: 1,
			0:  0,
			1:  1,
			2:  1,
		},
		"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99": map[int]int{
			7: 999,
			8: 1000,
			9: 1001,
		},
	}

	for prog, expected := range tests {
		for input, expectedOutput := range expected {
			t.Run(fmt.Sprintf("Day 2 - %s - %d", prog, input), func(t *testing.T) {
				inputs := make(chan int, 1)
				inputs <- input
				outputs := make(chan int)
				var output int
				wg := sync.WaitGroup{}
				wg.Add(1)
				go func() {
					for output = range outputs {
						t.Logf("Got output: %d", output)
					}
					wg.Done()
				}()
				debug = t.Logf
				t.Logf("Loading...\nprog:\t%s\ninput:\t%d\n", prog, input)

				mach := newMachine(prog, inputs, outputs)
				assert.NotPanics(t, func() {
					mach.run()
				})
				close(inputs)
				close(outputs)
				wg.Wait()
				t.Logf("Final state:\n\t%s\n", mach.String())
				assert.Equal(t, expectedOutput, output)
			})
		}
	}
}

func TestPositionMode(t *testing.T) {
	assert.Equal(t, 0, paramMode(10, 0))
	assert.Equal(t, 1, paramMode(10, 1))
}

func TestSequence(t *testing.T) {
	type sequenceTest struct {
		program       string
		phaseSequence []int
		result        int
	}

	tests := []sequenceTest{
		sequenceTest{
			program:       "",
			phaseSequence: []int{},
			result:        0,
		},
		sequenceTest{
			program:       "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			phaseSequence: []int{4, 3, 2, 1, 0},
			result:        43210,
		},
		sequenceTest{
			program:       "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			phaseSequence: []int{0, 1, 2, 3, 4},
			result:        54321,
		},
		sequenceTest{
			program:       "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			phaseSequence: []int{1, 0, 4, 3, 2},
			result:        65210,
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Day 7 check sequence %d", id), func(t *testing.T) {
			// debug = t.Logf
			t.Logf("Loading...\nprog:\t%s\ninputs:\t%#v\n", test.program, test.phaseSequence)

			var result int
			assert.NotPanics(t, func() {
				result = testPhaseSequence(test.phaseSequence, test.program)
			})
			assert.Equal(t, test.result, result)
		})
	}
}

func TestPart1Answer(t *testing.T) {
	debug = noOut
	assert.Equal(t, 95757, part1())
}

func TestPart2Answer(t *testing.T) {
	debug = noOut
	assert.Equal(t, 4275738, part2())
}

func TestMainRuns(t *testing.T) {
	assert.NotPanics(t, func() { main() })
}
