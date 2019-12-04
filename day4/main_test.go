package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Examples(t *testing.T) {
	goodPasswords := []int{
		122345,
		111111,
		111123,
	}
	badPasswords := []int{
		12234,
		223450,
		123789,
		135679,
	}
	for _, good := range goodPasswords {
		assert.True(t, validatePass(good, false), fmt.Sprintf("%d should be valid", good))
	}

	for _, bad := range badPasswords {
		assert.False(t, validatePass(bad, false), fmt.Sprintf("%d should be invalid", bad))
	}
}

func TestDay2Examples(t *testing.T) {
	goodPasswords := []int{
		122345,
	}
	badPasswords := []int{
		223450,
		123789,
		135679,
		111123,
		111111,
		123444,
	}
	for _, good := range goodPasswords {
		assert.True(t, validatePass(good, true), fmt.Sprintf("%d should be valid", good))
	}

	for _, bad := range badPasswords {
		assert.False(t, validatePass(bad, true), fmt.Sprintf("%d should be invalid", bad))
	}
}

func TestAnswers(t *testing.T) {
	assert.Equal(t, day1(), 921)
	assert.Equal(t, day2(), 603)
}

func TestMainRuns(t *testing.T) {
	assert.NotPanics(t, func() { main() })
}
