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

type point struct {
	x, y int
}

type grid map[point]tile

type game struct {
	lock           *sync.Mutex
	tiles          grid
	displayHistory []string
	maxX, maxY     int
	ballX          int
	paddleX        int
}

func (s *game) String() string {
	newOutput := ""
	for y := 0; y <= s.maxY; y++ {
		for x := 0; x <= s.maxX; x++ {
			switch s.tiles[point{x, y}] {
			case tileEmpty:
				newOutput += fmt.Sprint(" ")
			case tileWall:
				newOutput += fmt.Sprint("#")
			case tileBlock:
				newOutput += fmt.Sprint("~")
			case tilePaddle:
				newOutput += fmt.Sprint("-")
			case tileBall:
				newOutput += fmt.Sprint("o")
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
	s.tiles[point{x, y}] = tileType
}

func (s *game) autopilot() int64 {
	s.lock.Lock()
	ball := s.ballX
	paddle := s.paddleX
	s.lock.Unlock()
	if ball < paddle {
		return -1
	} else if ball > paddle {
		return 1
	} else {
		return 0
	}
}

func (s *game) outputCountHandler(wg *sync.WaitGroup, output chan int64, blockTileCount *int) {
	for range output {
		s.lock.Lock()
		<-output
		tileType := tile(<-output)

		if tile(tileType) == tileBlock {
			*blockTileCount++
		}
		s.lock.Unlock()
	}
	wg.Done()
}

func (s *game) outputHandler(wg *sync.WaitGroup, output chan int64, score *int) {
	for x := range output {
		s.lock.Lock()
		y := <-output
		tileType := tile(<-output)

		switch tileType {
		case tileBall:
			s.ballX = int(x)
		case tilePaddle:
			s.paddleX = int(x)
		}
		s.lock.Unlock()
		if x == -1 && y == 0 {
			*score = int(tileType)
		}
	}
	wg.Done()
}

func runGame(program string, play bool) int {
	score := 0
	gameInst := game{
		lock:  &sync.Mutex{},
		tiles: grid{},
	}
	blockTileCount := 0
	output := make(chan int64)
	wg := sync.WaitGroup{}

	cabinet := newMachine(program, nil, output)

	wg.Add(1)
	if play {
		cabinet.inputCallback = gameInst.autopilot
		cabinet.values[0] = 2
		go gameInst.outputHandler(&wg, output, &score)
	} else {
		go gameInst.outputCountHandler(&wg, output, &blockTileCount)
	}
	cabinet.run()
	wg.Wait()

	if play {
		return score
	}
	return blockTileCount
}
