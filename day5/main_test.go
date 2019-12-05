package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Examples(t *testing.T) {
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

	for input, expected := range tests {
		mach := newMachine(input, 1002)
		mach.run()
		assert.Equal(t, expected, mach.String())
	}
}

func TestDay2Examples(t *testing.T) {
	type day2inputtest map[int]int

	tests := map[string]day2inputtest{
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
		for input, output := range expected {
			t.Run(fmt.Sprintf("Day 2 - %s - %d", prog, input), func(t *testing.T) {
				debug = t.Logf
				t.Logf("Loading...\nprog:\t%s\ninput:\t%d\n", prog, input)

				mach := newMachine(prog, input)
				assert.NotPanics(t, func() {
					mach.run()
				})
				t.Logf("Final state:\n\t%s\n", mach.String())
				assert.Equal(t, output, mach.lastOutput)
			})
		}
	}
}

func TestPositionMode(t *testing.T) {
	debug = noOut
	mach := newMachine("99", 1)
	assert.Equal(t, 0, mach.paramMode(10, 0))
	assert.Equal(t, 1, mach.paramMode(10, 1))
}

func TestDay1Answer(t *testing.T) {
	debug = noOut
	assert.Equal(t, 15097178, day1())
}

func TestDay2Answer(t *testing.T) {
	debug = noOut
	assert.Equal(t, 1558663, day2())
}
