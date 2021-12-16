package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	versionSum, value, err := parse(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
	}
	return versionSum, value
}

func parse(in string) (int, int, error) {
	packetBitBuilder := strings.Builder{}
	for _, ch := range []byte(in) {
		val, err := strconv.ParseUint(string(ch), 16, 8)
		if err != nil {
			return -1, -1, err
		}
		packetBitBuilder.WriteString(fmt.Sprintf("%04b", val))
	}
	packetBits := packetBitBuilder.String()
	sum, value, _, err := parseBits(packetBits, 0, 0)
	return sum, value, err
}

func parseBits(packetBits string, maxPackets int, operation packetType) (int, int, int, error) {
	versionSum := 0
	packetValues := []int{}
	packetValue := 0
	packetsRead := 0
	idx := 0

	for strings.Contains(packetBits[idx:], "1") {
		packetsRead++
		if maxPackets > 0 && packetsRead > maxPackets {
			break
		}
		version, err := strconv.ParseInt(packetBits[idx:idx+3], 2, 8)
		if err != nil {
			return versionSum, packetValue, idx, err
		}
		versionSum += int(version)
		idx += 3

		pType, err := strconv.ParseInt(packetBits[idx:idx+3], 2, 8)
		if err != nil {
			return versionSum, packetValue, idx, err
		}
		idx += 3

		switch packetType(pType) {
		case packetTypeLiteral:
			litString := strings.Builder{}
			for {
				group := packetBits[idx : idx+5]
				idx += 5
				litString.WriteString(group[1:])
				if group[0] == '0' {
					break
				}
			}
			value, err := strconv.ParseInt(litString.String(), 2, 63)
			if err != nil {
				return versionSum, packetValue, idx, err
			}
			packetValues = append(packetValues, int(value))
		default:
			lengthType := packetBits[idx]
			idx++
			if lengthType == '0' {
				length, err := strconv.ParseInt(packetBits[idx:idx+15], 2, 32)
				idx += 15
				if err != nil {
					return versionSum, packetValue, idx, err
				}
				subPacket := packetBits[idx : idx+int(length)]
				idx += int(length)
				subSum, subValue, consumed, err := parseBits(subPacket, 0, packetType(pType))
				packetValues = append(packetValues, subValue)
				versionSum += subSum
				if err != nil {
					return versionSum, packetValue, idx, err
				}
				if consumed != int(length) {
					fmt.Println("Wrong number of bits comsumed?", consumed, length)
				}

			} else {
				length, err := strconv.ParseUint(packetBits[idx:idx+11], 2, 11)
				idx += 11
				if err != nil {
					return versionSum, packetValue, idx, err
				}
				subSum, subValue, consumed, err := parseBits(packetBits[idx:], int(length), packetType(pType))
				packetValues = append(packetValues, subValue)
				versionSum += subSum
				idx += consumed
				if err != nil {
					return versionSum, packetValue, idx, err
				}
			}
		}
	}
	switch packetType(operation) {
	case packetTypeSum:
		packetValue = 0
		for _, val := range packetValues {
			packetValue += val
		}
	case packetTypeProduct:
		packetValue = 1
		for _, val := range packetValues {
			packetValue *= val
		}
	case packetTypeMin:
		packetValue = utils.MaxInt
		for _, val := range packetValues {
			if val < packetValue {
				packetValue = val
			}
		}
	case packetTypeMax:
		packetValue = 0
		for _, val := range packetValues {
			if val > packetValue {
				packetValue = val
			}
		}
	case packetTypeLiteral:
		packetValue = packetValues[0]
	case packetTypeGreater:
		if packetValues[0] > packetValues[1] {
			packetValue = 1
		} else {
			packetValue = 0
		}
	case packetTypeLess:
		if packetValues[0] < packetValues[1] {
			packetValue = 1
		} else {
			packetValue = 0
		}
	case packetTypeEqual:
		if packetValues[0] == packetValues[1] {
			packetValue = 1
		} else {
			packetValue = 0
		}
	default:
		return versionSum, packetValue, idx, fmt.Errorf("Unhandled operation %d", operation)
	}
	return versionSum, packetValue, idx, nil
}

type packetType int

const (
	packetTypeSum packetType = iota
	packetTypeProduct
	packetTypeMin
	packetTypeMax
	packetTypeLiteral
	packetTypeGreater
	packetTypeLess
	packetTypeEqual
)

var packetTypeName = map[packetType]string{
	packetTypeSum:     "sum",
	packetTypeProduct: "product",
	packetTypeMin:     "min",
	packetTypeMax:     "max",
	packetTypeLiteral: "literal",
	packetTypeGreater: "greater",
	packetTypeLess:    "less",
	packetTypeEqual:   "equal",
}

var benchmark = false
