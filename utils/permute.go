package utils

// Implementation of Heaps algorithm:
//   https://en.wikipedia.org/wiki/Heap%27s_algorithm
func PermuteInts(input []int) [][]int {
	output := [][]int{}

	var generator func(int, []int)
	generator = func(k int, A []int) {
		if k == 1 {
			tmp := make([]int, len(A))
			copy(tmp, A)
			output = append(output, tmp)
		} else {
			for i := 0; i < k; i++ {
				generator(k-1, A)
				if k%2 == 1 {
					tmp := A[i]
					A[i] = A[k-1]
					A[k-1] = tmp
				} else {
					tmp := A[0]
					A[0] = A[k-1]
					A[k-1] = tmp
				}
			}
		}
	}
	generator(len(input), input)
	return output
}
