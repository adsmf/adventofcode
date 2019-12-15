package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/adsmf/adventofcode2019/utils/intcode"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var interactive bool
var autopilot bool

func init() {
	flag.BoolVar(&interactive, "interactive", false, "Run game interactively")
	flag.BoolVar(&autopilot, "autopilot", false, "Turn on autopilot for interactive mode")
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
	tiles           grid
	displayHistory  []string
	maxX, maxY      int
	ballX           int
	paddleX         int
	lastDraw        string
	screen          *tview.TextView
	cabinet         *tview.Application
	nextInput       int
	cpu             *intcode.Machine
	piloter         intcode.InputCallback
	bufferedOutputs []int
	score           int
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

func (s *game) pilot() (int, bool) {
	return s.piloter()
}

func (s *game) autopilot() (int, bool) {
	ball := s.ballX
	paddle := s.paddleX
	if interactive {
		time.Sleep(10 * time.Millisecond)
	}
	if ball < paddle {
		return -1, false
	} else if ball > paddle {
		return 1, false
	} else {
		return 0, false
	}
}

func (s *game) manualpilot() (int, bool) {
	time.Sleep(1 * time.Second)
	returnValue := s.nextInput
	s.nextInput = 0
	return returnValue, false
}

type blockCounter struct {
	opCount        int
	blockTileCount int
}

func (s *blockCounter) outputCountHandler(output int) {
	s.opCount++
	if s.opCount == 3 {
		tileType := tile(output)
		if tile(tileType) == tileBlock {
			s.blockTileCount++
		}
		s.opCount = 0
	}
}

func (s *game) outputHandler(output int) {
	s.bufferedOutputs = append(s.bufferedOutputs, output)
	if len(s.bufferedOutputs) < 3 {
		return
	}
	x := s.bufferedOutputs[0]
	y := s.bufferedOutputs[1]
	tileType := tile(s.bufferedOutputs[2])
	s.bufferedOutputs = []int{}

	switch tileType {
	case tileBall:
		s.ballX = x
	case tilePaddle:
		s.paddleX = x
	}
	if x == -1 && y == 0 {
		s.score = int(tileType)
	}
	s.set(x, y, tileType)
	lastDraw := s.String()
	if s.screen != nil {
		s.cabinet.QueueUpdateDraw(func() {
			s.screen.SetText(fmt.Sprintf("Ball: %d; Paddle: %d\n%s", s.ballX, s.paddleX, lastDraw))
		})
	}
}

func (s *game) keyboadHandler(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter:
		if autopilot {
			s.piloter = s.manualpilot
		} else {
			s.piloter = s.autopilot
		}
		autopilot = !autopilot
	case tcell.KeyLeft, ',', 'a':
		s.nextInput = -1
	case tcell.KeyRight, '.', 'd':
		s.nextInput = 1
	case 'q':
		s.cabinet.Stop()
		os.Exit(0)
	}
	return event
}

func runGame(program string, play bool, interactive bool) int {
	gameInst := game{
		tiles:   grid{},
		piloter: nil,
	}
	var opHandler intcode.OutputCallback
	counter := blockCounter{}
	if play {
		opHandler = gameInst.outputHandler
	} else {
		opHandler = counter.outputCountHandler
	}
	cpu := intcode.NewMachine(intcode.M19(gameInst.pilot, opHandler))
	cpu.LoadProgram(program)
	gameInst.cpu = &cpu

	if play && interactive {
		if autopilot {
			gameInst.piloter = gameInst.autopilot
		} else {
			gameInst.piloter = gameInst.manualpilot
		}
		cpu.WriteRAM(0, 2)
		mainView := tview.NewTextView()
		mainView.SetBorder(true).SetTitle("Int(eractive)")

		gameInst.cabinet = tview.NewApplication().SetRoot(mainView, true)
		gameInst.screen = mainView

		gameInst.cabinet.SetInputCapture(gameInst.keyboadHandler)
	} else if play {
		gameInst.piloter = gameInst.autopilot
		cpu.WriteRAM(0, 2)
	}
	if gameInst.cabinet != nil {
		go gameInst.cabinet.Run()
	}
	cpu.Run(false)

	if gameInst.cabinet != nil {
		wg := sync.WaitGroup{}
		wg.Add(1)
		modal := tview.NewModal()
		modal.
			SetText(fmt.Sprintf("Score: %d", gameInst.score)).
			AddButtons([]string{"Quit"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				wg.Done()
				gameInst.cabinet.Stop()
				os.Exit(0)
			}).
			SetTitle("~ FIN ~")
		gameInst.cabinet.SetRoot(modal, false)
		wg.Wait()
	}

	if play {
		return gameInst.score
	}
	return counter.blockTileCount
}

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}
