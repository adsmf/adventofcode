package main

import (
	"fmt"
	"github.com/adsmf/adventofcode2019/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	img := parseImage(utils.ReadInputLines("input.txt")[0], 25, 6)
	l := findLeastZeros(img)
	return oneTimesTwo(l)
}

func part2() int {
	img := parseImage(utils.ReadInputLines("input.txt")[0], 25, 6)
	check := render(img, 25, 6)
	return check
}

func render(img image, w, h int) int {

	for layerNum, l := range img {
		if layerNum == 0 {
			continue
		}
		for pos, p := range l {
			if img[0][pos] == 2 {
				img[0][pos] = p
			}
		}
	}
	checksum := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if img[0][y*w+x] == 1 {
				fmt.Print("#")
				checksum++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return checksum
}

func oneTimesTwo(l layer) int {
	ones := 0
	twos := 0
	for _, p := range l {
		if p == 1 {
			ones++
		}
		if p == 2 {
			twos++
		}
	}
	return ones * twos
}

func findLeastZeros(img image) layer {
	bestLayer := layer{}
	bestCount := utils.MaxInt
	for _, layer := range img {
		count := 0
		for _, p := range layer {
			if p == 0 {
				count++
			}
		}
		if count < bestCount {
			bestCount = count
			bestLayer = layer
		}
	}
	return bestLayer
}

type pixel byte
type layer []pixel
type image []layer

func parseImage(data string, w, h int) image {
	layerSize := w * h
	numLayers := len(data) / layerSize
	img := make(image, numLayers)
	for pos, char := range data {
		layerNum := pos / layerSize
		layerPos := pos % layerSize
		if img[layerNum] == nil {
			img[layerNum] = make(layer, layerSize)
		}
		img[layerNum][layerPos] = pixel(char - '0')
	}
	return img
}
