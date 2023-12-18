package leetcode_989

func CaseTwo(num []int, k int) []int {
	var total = make([]int, 0)
	for l := len(num) - 1; l >= 0 || k > 0; l-- {
		if l >= 0 {
			k += num[l]
		}

		total = append([]int{k % 10}, total...)
		k /= 10
	}

	return total
}
