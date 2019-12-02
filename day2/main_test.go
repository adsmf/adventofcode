package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Examples(t *testing.T) {
	tests := map[string]string{
		"1,0,0,0,99":                    "2,0,0,0,99",
		"2,3,0,3,99":                    "2,3,0,6,99",
		"2,4,4,5,99,0":                  "2,4,4,5,99,9801",
		"1,1,1,4,99,5,6,0,99":           "30,1,1,4,2,5,6,0,99",
		"1,9,10,3,2,3,11,0,99,30,40,50": "3500,9,10,70,2,3,11,0,99,30,40,50",
	}

	for input, expected := range tests {
		mach := newMachine(input)
		mach.run()
		assert.Equal(t, expected, mach.String())
	}
}

func TestDay1Answer(t *testing.T) {
	assert.Equal(t, 3224742, day1())
}

func TestDay2Answer(t *testing.T) {
	assert.Equal(t, 7960, day2())
}
