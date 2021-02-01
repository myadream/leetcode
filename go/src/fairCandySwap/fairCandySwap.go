package fairCandySwap

import (
	"sort"
)

//https://leetcode-cn.com/problems/fair-candy-swap/
func fairCandySwap(A []int, B []int) []int {
	var dis, sum, start_A, start_B int

	for _, v := range A {
		sum += v
	}

	for _, v := range B {
		sum -= v
	}

	dis = sum / 2

	sort.Ints(A)
	sort.Ints(B)

	for true {
		if A[start_A] - B[start_B] == dis {
			return []int{A[start_A], B[start_B]}
		} else if A[start_A] - B[start_B] > dis {
			start_B += 1
		} else {
			start_A += 1
		}
	}

	return []int{}
}