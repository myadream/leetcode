package template

import "math"

func CaseOne(heights []int) []int {
	sums := make([]int, len(heights))
	stack := []int{math.MaxInt}

	for i := len(heights) - 1; i >= 0; i-- {
		for stack[len(stack)-1] <= heights[i] {
			stack = stack[:len(stack)-1]
			sums[i] += 1
		}

		if len(stack) > 1 {
			sums[i] += 1
		}

		stack = append(stack, heights[i])
	}

	return sums
}
