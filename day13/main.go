package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	inputString := loadInputString()
	return runGame(inputString, false)
}

func part2() int {
	inputString := loadInputString()
	return runGame(inputString, true)
}

type tile int

const (
	tileEmpty  tile = iota // 0 is an empty tile. No game object appears in this tile.
	tileWall               // 1 is a wall tile. Walls are indestructible barriers.
	tileBlock              // 2 is a block tile. Blocks can be broken by the ball.
	tilePaddle             // 3 is a horizontal paddle tile. The paddle is indestructible.
	tileBall               // 4 is a ball tile. The ball moves diagonally and bounces off objects.
)

type game struct {
	lock           *sync.Mutex
	tiles          map[int]map[int]tile
	displayHistory []string
	maxX, maxY     int
	ballX          int
	paddleX        int
}

func (s *game) String() string {
	newOutput := ""
	for y := 0; y <= s.maxY; y++ {
		for x := 0; x <= s.maxX; x++ {
			switch s.tiles[x][y] {
			case tileEmpty:
				newOutput += fmt.Sprint(" ")
			case tileWall:
				newOutput += fmt.Sprint("#")
			case tileBlock:
				newOutput += fmt.Sprint("~")
			case tilePaddle:
				newOutput += fmt.Sprint("-")
				s.paddleX = x
			case tileBall:
				newOutput += fmt.Sprint("o")
				s.ballX = x
			default:
				newOutput += fmt.Sprint("?")
			}
		}
		newOutput += fmt.Sprintln()
	}
	return newOutput
}

func (s *game) set(x, y int, tileType tile) {
	if x > s.maxX {
		s.maxX = x
	}
	if y > s.maxY {
		s.maxY = y
	}
	if s.tiles == nil {
		s.tiles = map[int]map[int]tile{}
	}
	if s.tiles[x] == nil {
		s.tiles[x] = map[int]tile{}
	}
	switch tileType {
	case tileBall:
		s.ballX = x
	case tilePaddle:
		s.paddleX = x
	}
	s.tiles[x][y] = tileType
}

func runGame(program string, play bool) int {
	score := 0
	gameInst := game{lock: &sync.Mutex{}}
	blockTileCount := 0
	output := make(chan int64)
	input := make(chan int64, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for x := range output {
			gameInst.lock.Lock()
			y := <-output
			tileType := tile(<-output)

			if play == false {
				if tile(tileType) == tileBlock {
					blockTileCount++
				}
			}
			if x == -1 && y == 0 {
				score = int(tileType)
			}
			gameInst.set(int(x), int(y), tileType)
			gameInst.lock.Unlock()
		}
		wg.Done()
	}()

	cabinet := newMachine(program, input, output)
	cabinet.inputCallback = func() int64 {
		gameInst.lock.Lock()
		var direction int64
		ball := gameInst.ballX
		paddle := gameInst.paddleX
		if ball < paddle {
			direction = -1
		} else if ball > paddle {
			direction = 1
		} else {
			direction = 0
		}
		gameInst.lock.Unlock()
		return direction
	}
	if play {
		cabinet.values[0] = 2
	}
	cabinet.run()

	wg.Wait()
	if play {
		return score
	}
	return blockTileCount
}
