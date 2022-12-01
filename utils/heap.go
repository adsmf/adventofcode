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
	if len(h.heap) < h.max {
		h.heap = append(h.heap, val)
		return
	}
	min := h.heap[h.max-1]
	if val > min {
		for i := 0; i < h.max; i++ {
			if h.heap[i] < val {
				for j := h.max - 1; j > i; j-- {
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
