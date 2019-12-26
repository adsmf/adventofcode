package main

import (
	"fmt"

	"github.com/adsmf/adventofcode2019/utils"
	"github.com/adsmf/adventofcode2019/utils/intcode"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	n := newNetwork()
	return n.runTo255()
}

func part2() int {
	n := newNetwork()
	return n.double0()
}

type network struct {
	devices map[address]*nic

	currentNAT *packet
	lastNAT    *packet
	packetSent bool
}

func (n *network) double0() int {
	idleSteps := 0
	for steps := 0; ; steps++ {
		n.packetSent = false
		for i := 0; i < 50; i++ {
			n.devices[address(i)].machine.Step()
		}
		if n.packetSent {
			idleSteps = 0
		} else {
			idleSteps++

			if idleSteps > 4000 && n.currentNAT != nil {
				if n.lastNAT != nil && *n.lastNAT == *n.currentNAT {
					break
				}
				n.lastNAT = n.currentNAT
				n.send(0, *n.currentNAT)
				idleSteps = 0
			}
		}
	}
	return n.currentNAT.y
}

func (n *network) runTo255() int {
	for steps := 0; steps < 5000; steps++ {
		if n.currentNAT != nil {
			break
		}
		for i := 0; i < 50; i++ {
			n.devices[address(i)].machine.Step()
		}
	}
	return n.currentNAT.y
}

func (n *network) send(addr address, p packet) {
	if addr == 255 {
		n.currentNAT = &p
	} else {
		n.packetSent = true
		n.devices[addr].receive(p)
	}
}

func newNetwork() *network {
	n := network{
		devices: map[address]*nic{},
	}

	prog := utils.ReadInputLines("input.txt")[0]
	for i := 0; i < 50; i++ {
		addr := address(i)
		dev := nic{
			address:       addr,
			network:       &n,
			sendBuffer:    []int{i},
			receiveBuffer: []int{},
		}
		dev.machine = intcode.NewMachine(intcode.M19(dev.inputHandler, dev.outputHandler))
		err := dev.machine.LoadProgram(prog)
		if err != nil {
			panic(err)
		}

		n.devices[addr] = &dev
	}
	return &n
}

type address int

type packet struct {
	x, y int
}

type nic struct {
	address address
	machine intcode.Machine
	network *network

	sendBuffer    []int
	receiveBuffer []int
}

func (n *nic) receive(p packet) {
	n.sendBuffer = append(n.sendBuffer, p.x, p.y)
}

func (n *nic) inputHandler() (int, bool) {
	if len(n.sendBuffer) > 0 {
		var toSend int
		toSend, n.sendBuffer = n.sendBuffer[0], n.sendBuffer[1:]
		return toSend, false
	}
	return -1, false
}

func (n *nic) outputHandler(input int) {
	n.receiveBuffer = append(n.receiveBuffer, input)

	if len(n.receiveBuffer) >= 3 {
		addr, x, y := n.receiveBuffer[0], n.receiveBuffer[1], n.receiveBuffer[2]
		n.receiveBuffer = n.receiveBuffer[3:]

		n.network.send(address(addr), packet{x, y})
	}
}
