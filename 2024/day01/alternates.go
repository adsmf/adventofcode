package main

import (
	"sort"

	"github.com/adsmf/adventofcode/utils"
)

func solveInitial() (int, int) {
	nums := utils.GetInts(input)
	l1, l2 := make([]int, 0, len(nums)/2), make([]int, 0, len(nums)/2)
	counts2 := map[int]int{}
	for i := 0; i < len(nums); i += 2 {
		l1 = append(l1, nums[i])
		l2 = append(l2, nums[i+1])
		counts2[nums[i+1]]++
	}
	sort.Ints(l1)
	sort.Ints(l2)
	p1, p2 := 0, 0
	for i := range len(l1) {
		if l1[i] > l2[i] {
			p1 += l1[i] - l2[i]
		} else {
			p1 += l2[i] - l1[i]
		}
		p2 += counts2[l1[i]] * l1[i]
	}
	return p1, p2
}

func solveMapless() (int, int) {
	nums := utils.GetInts(input)
	l1, l2 := make([]int, 0, len(nums)/2), make([]int, 0, len(nums)/2)
	for i := 0; i < len(nums); i += 2 {
		l1 = append(l1, nums[i])
		l2 = append(l2, nums[i+1])
	}
	sort.Ints(l1)
	sort.Ints(l2)
	p1, p2 := 0, 0
	l2ptr := 0
	for i := range len(l1) {
		n1, n2 := l1[i], l2[i]
		if n1 > n2 {
			p1 += n1 - n2
		} else {
			p1 += n2 - n1
		}
		for ; l2ptr < len(l2) && l2[l2ptr] < n1; l2ptr++ {
		}
		for ; l2ptr < len(l2) && l2[l2ptr] == n1; l2ptr++ {
			p2 += n1
		}
	}
	return p1, p2
}

func solveSingleSlice() (int, int) {
	nums := utils.GetInts(input)
	offset := len(nums) / 2
	for i := 1; i < offset; i++ {
		nums[i], nums[i*2] = nums[i*2], nums[i]
	}
	l1 := nums[:offset]
	l2 := nums[offset:]
	sort.Ints(l1)
	sort.Ints(l2)
	p1, p2 := 0, 0
	l2ptr := 0
	for i := range len(l1) {
		n1, n2 := l1[i], l2[i]
		if n1 > n2 {
			p1 += n1 - n2
		} else {
			p1 += n2 - n1
		}
		for ; l2ptr < len(l2) && l2[l2ptr] < n1; l2ptr++ {
		}
		for ; l2ptr < len(l2) && l2[l2ptr] == n1; l2ptr++ {
			p2 += n1
		}
	}
	return p1, p2
}
