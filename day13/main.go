package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/rivo/tview"
)

var interactive bool

func init() {
	flag.BoolVar(&interactive, "interactive", false, "Run game interactively")
}

func main() {
	flag.Parse()
	if interactive {
		inputString := loadInputString()
		runGame(inputString, true, true)
	} else {
		fmt.Printf("Part 1: %d\n", part1())
		fmt.Printf("Part 2: %d\n", part2())
	}
}

func part1() int {
	inputString := loadInputString()
	return runGame(inputString, false, false)
}

func part2() int {
	inputString := loadInputString()
	return runGame(inputString, true, false)
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
	lastDraw       string
	screen         *tview.TextView
	cabinet        *tview.Application
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
		if x == -1 && y == 0 {
			*score = int(tileType)
		}
		s.set(int(x), int(y), tileType)
		lastDraw := s.String()
		s.lock.Unlock()
		if s.screen != nil {
			s.cabinet.QueueUpdateDraw(func() {
				s.screen.SetText(fmt.Sprintf("Ball: %d; Paddle: %d\n%s", s.ballX, s.paddleX, lastDraw))
				// s.cabinet.Draw()
			})
		}
	}
	if s.cabinet != nil {
		modal := tview.NewModal()
		modal.
			SetText(fmt.Sprintf("Score: %d", *score)).
			AddButtons([]string{"Quit"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				s.cabinet.Stop()
			}).
			SetTitle("~ FIN ~")
		s.cabinet.SetRoot(modal, false)
	}
	wg.Done()
}

func runGame(program string, play bool, interactive bool) int {
	score := 0
	gameInst := game{
		lock:  &sync.Mutex{},
		tiles: grid{},
	}
	blockTileCount := 0
	output := make(chan int64)
	wg := sync.WaitGroup{}

	cpu := newMachine(program, nil, output)

	wg.Add(1)
	if play && interactive {
		cpu.inputCallback = gameInst.autopilot
		cpu.values[0] = 2
		mainView := tview.NewTextView()
		mainView.SetBorder(true).SetTitle("Int(eractive)")

		gameInst.cabinet = tview.NewApplication().SetRoot(mainView, true)
		gameInst.screen = mainView

		go gameInst.outputHandler(&wg, output, &score)
	} else if play {
		cpu.inputCallback = gameInst.autopilot
		cpu.values[0] = 2
		go gameInst.outputHandler(&wg, output, &score)
	} else {
		go gameInst.outputCountHandler(&wg, output, &blockTileCount)
	}
	go cpu.run()
	if gameInst.cabinet != nil {
		gameInst.cabinet.Run()
	}
	wg.Wait()

	if play {
		return score
	}
	return blockTileCount
}
