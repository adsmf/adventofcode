package utils

import (
	"golang.org/x/exp/constraints"
)

type TopN[T constraints.Ordered] struct {
	heap []T
	max  int
}

func NewTopN[T constraints.Ordered](max int) TopN[T] {
	return TopN[T]{
		heap: make([]T, 0, max),
		max:  max,
	}
}

func (h *TopN[T]) Add(val T) {
	min := (T)(0)
	if len(h.heap) < h.max {
		h.heap = append(h.heap, (T)(0))
	} else if len(h.heap) > 0 {
		min = h.heap[len(h.heap)-1]
	}
	if val > min {
		for i := 0; i < len(h.heap); i++ {
			if h.heap[i] < val {
				for j := len(h.heap) - 1; j > i; j-- {
					h.heap[j] = h.heap[j-1]
				}
				h.heap[i] = val
				return
			}
		}
	}
}

func (h *TopN[T]) Values() []T {
	return h.heap
}
