package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcLevel(t *testing.T) {
	type exampleLevel struct {
		x, y          int
		serial        int
		expectedLevel int
	}
	examples := []exampleLevel{
		exampleLevel{3, 5, 8, 4},
		exampleLevel{122, 79, 57, -5},
		exampleLevel{217, 196, 39, 0},
		exampleLevel{101, 153, 71, 4},
	}
	for exampleID, example := range examples {
		t.Run(fmt.Sprintf("Example %d", exampleID), func(t *testing.T) {
			t.Logf("Running example %+v", example)
			level := calcLevel(example.serial, example.x, example.y)
			assert.Equal(t, example.expectedLevel, level)
		})
	}
}

func TestBestGrid(t *testing.T) {
	type exampleGrid struct {
		serial int
		x, y   int
		size   int
		level  int
	}
	examples := []exampleGrid{
		exampleGrid{18, 33, 45, 3, 29},
		exampleGrid{42, 21, 61, 3, 30},
		exampleGrid{18, 90, 269, 16, 113},
		exampleGrid{42, 232, 251, 12, 119},
	}
	for exampleID, example := range examples {
		t.Run(fmt.Sprintf("Example %d", exampleID), func(t *testing.T) {
			t.Logf("Running example %+v", example)
			level, x, y := bestForSerial(example.serial, example.size)
			assert.Equal(t, example.level, level)
			assert.Equal(t, example.x, x)
			assert.Equal(t, example.y, y)
		})
	}
}

func TestBestGridSize(t *testing.T) {
	type exampleGrid struct {
		serial int
		x, y   int
		size   int
		level  int
	}
	examples := []exampleGrid{
		exampleGrid{18, 90, 269, 16, 113},
		exampleGrid{42, 232, 251, 12, 119},
	}
	for exampleID, example := range examples {
		t.Run(fmt.Sprintf("Example %d", exampleID), func(t *testing.T) {
			t.Logf("Running example %+v", example)
			level, x, y, size := bestGrid(example.serial)
			assert.Equal(t, example.level, level)
			assert.Equal(t, example.x, x)
			assert.Equal(t, example.y, y)
			assert.Equal(t, example.size, size)
		})
	}
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 20,54
	//Part 2: 233,93,13
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
