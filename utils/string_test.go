package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInts(t *testing.T) {
	tests := []struct {
		input        string
		expectedInts []int
	}{
		{"1 2 5 6", []int{1, 2, 5, 6}},
		{"1 2\n5 6", []int{1, 2, 5, 6}},
		{"28,93,0,50,65,87", []int{28, 93, 0, 50, 65, 87}},
		{"1946127596-1953926346", []int{1946127596, 1953926346}},
		{"0-1888888", []int{0, 1888888}},
	}
	for _, test := range tests {
		t.Run("test", func(t *testing.T) {
			ints := GetInts(test.input)
			assert.EqualValues(t, test.expectedInts, ints)
		})
	}
}
