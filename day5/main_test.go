package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Examples(t *testing.T) {
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
		// "3,1,99": "3,1002,99",
	}

	for input, expected := range tests {
		mach := newMachine(input, 1002)
		mach.run()
		assert.Equal(t, expected, mach.String())
	}
}

func TestPositionMode(t *testing.T) {
	mach := newMachine("99", 1)
	assert.Equal(t, 0, mach.paramMode(10, 0))
	assert.Equal(t, 1, mach.paramMode(10, 1))
}

func TestDay1Answer(t *testing.T) {
	assert.Equal(t, 15097178, day1())
}

func TestDay2Answer(t *testing.T) {
	assert.Equal(t, 0, day2())
}
