package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/adsmf/adventofcode2019/utils"
	"github.com/adsmf/adventofcode2019/utils/intcode"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	a := newAdventure()
	a.run()
}

type adventure struct {
	app        *tview.Application
	outputView *tview.TextView
	inputBox   *tview.InputField

	log       string
	capturing bool
	sideLog   string

	command  string
	lastChar byte

	game intcode.Machine
}

func newAdventure() *adventure {
	a := &adventure{}

	prog := loadInput("input.txt")
	a.game = intcode.NewMachine(intcode.M19(a.inputHandler, a.outputHandler), intcode.DecodeOps())
	err := a.game.LoadProgram(prog)
	if err != nil {
		panic(err)
	}

	app := tview.NewApplication()
	a.app = app

	outputView := tview.NewTextView()
	a.outputView = outputView
	outputView.SetBorder(true).SetTitle("Int(eractive|code) Advent(ure| of code)")

	inputView := tview.NewForm()
	inputField := tview.NewInputField().
		SetLabel("Command: ").
		SetFieldWidth(50).
		SetDoneFunc(a.sendCommand)
	a.inputBox = inputField
	inputView.AddFormItem(inputField)

	mainLayout := tview.NewGrid().
		SetRows(0, 3).
		SetBorders(true)

	mainLayout.AddItem(outputView, 0, 0, 1, 1, 20, 0, false)
	mainLayout.AddItem(inputView, 1, 0, 1, 1, 1, 0, true)

	app.SetRoot(mainLayout, true)

	a.log = "The adventure begins...\n\n"
	return a
}

func (a *adventure) run() {
	north := func() { a.command += fmt.Sprintln("north") }
	south := func() { a.command += fmt.Sprintln("south") }
	east := func() { a.command += fmt.Sprintln("east") }
	west := func() { a.command += fmt.Sprintln("west") }
	take := func(obj item) { a.command += fmt.Sprintf("take %v\n", obj) }
	inv := func() { a.command += fmt.Sprintln("inv") }
	// Collect all the things (route discovered by playing manually)
	east()
	take(loom)
	east()
	take(fixedpoint)
	north()
	take(spoolofcat6)
	west()
	take(shell)
	east()
	north()
	take(weathermachine)
	south()
	south()
	west()
	south()
	take(ornament)
	west()
	north()
	take(candycane)
	south()
	east()
	north()
	west()
	north()
	north()
	east()
	// Finish at security checkpoint
	inv()

	go a.game.Run(false)

	go a.tryItems()
	err := a.app.Run()
	if err != nil {
		panic(err)
	}
}

type item int

const (
	nothing        item = 0
	loom           item = 1
	spoolofcat6    item = 2
	fixedpoint     item = 4
	weathermachine item = 8
	ornament       item = 16
	wreath         item = 32
	shell          item = 64
	candycane      item = 128
	END            item = 256
)

func (i item) String() string {
	switch i {
	case loom:
		return "loom"
	case spoolofcat6:
		return "spool of cat6"
	case fixedpoint:
		return "fixed point"
	case weathermachine:
		return "weather machine"
	case ornament:
		return "ornament"
	case wreath:
		return "wreath"
	case shell:
		return "shell"
	case candycane:
		return "candy cane"
	}
	return "?!?!???!?!?!?!?"
}

func (a *adventure) tryItems() {
	last := item(255)
	waitPrompt := func() {
		for !strings.Contains(a.log, "Command?") ||
			len(a.command) > 0 {
			time.Sleep(10 * time.Millisecond)
		}
		time.Sleep(10 * time.Millisecond)

		a.app.QueueUpdateDraw(func() {
			a.outputView.SetText(a.log)
			a.outputView.ScrollToEnd()
		})
	}
	waitPrompt()
	a.capturing = true
	for try := item(0); try < END; try++ {
		delta := try ^ last
		for bit := 0; bit < 8; bit++ {
			deltaObj := (delta & (1 << bit))
			if deltaObj > 0 {
				var action string
				if (try & (1 << bit)) > 0 {
					action = "take"
				} else {
					action = "drop"
				}
				a.command += fmt.Sprintf("%s %v\n", action, deltaObj)
				waitPrompt()
			}
		}

		a.sideLog = ""
		a.log += fmt.Sprintf("Try: %d\n", try)
		a.command += "south\n"
		sleeps := 0
		for !strings.Contains(a.sideLog, "Command?") ||
			len(a.command) > 0 {
			if sleeps > 1000 {
				a.log += "!command\n"
				break
			}
			time.Sleep(1 * time.Millisecond)
			sleeps++
		}
		if !strings.Contains(a.sideLog, "ejected") {
			a.log += "!EJECTED\n"
			break
		}
		last = try
	}
	a.log += a.sideLog
	a.capturing = false
	a.app.QueueUpdateDraw(func() {
		a.outputView.SetText(a.log)
		a.outputView.ScrollToEnd()
	})
}

func (a *adventure) inputHandler() (int, bool) {
	for {
		if len(a.command) > 0 {
			var char byte
			char, a.command = a.command[0], a.command[1:]
			if a.lastChar <= '\n' {
				a.log += "> "
			}
			a.lastChar = char
			a.log += fmt.Sprintf("%c", char)
			return int(char), false
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (a *adventure) sendCommand(key tcell.Key) {
	input := a.inputBox.GetText()
	a.inputBox.SetText("")
	a.command = input + fmt.Sprintf("%c", 10)
	a.app.QueueUpdateDraw(func() {
		a.outputView.SetText(a.log)
		a.outputView.ScrollToEnd()
	})
}

func (a *adventure) outputHandler(inp int) {
	if a.capturing {
		a.sideLog += fmt.Sprintf("%c", inp)
		return
	}
	if len(a.command) > 0 {
		return
	}
	a.log += fmt.Sprintf("%c", inp)
	a.app.QueueUpdateDraw(func() {
		a.outputView.SetText(a.log)
		a.outputView.ScrollToEnd()
	})
}

func loadInput(filename string) string {
	prog := utils.ReadInputLines(filename)[0]
	return prog
}
