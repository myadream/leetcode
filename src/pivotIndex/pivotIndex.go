package pivotIndex

import "fmt"

//https://leetcode-cn.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	var total, sum int
	var n = len(nums)

	for i := 0; i < n; i++ {
		total += nums[i]
	}

	for i := 0; i < n; i++ {
		fmt.Println(2*sum+nums[i], total)
		if 2*sum+nums[i] == total {
			return i
		}

		sum += nums[i]
	}

	return -1
}

//并查集... 待完善
// func pivotIndex(nums []int) int {
// 	var sum, n = 0, len(nums)
// 	var prefix = make([]int, n)

// 	for k, v := range nums {
// 		sum += v
// 		prefix[k] = sum
// 	}

// 	for k, _ := range nums {
// 		if prefix[k] == prefix[k]-prefix[k+1] {
// 			return k
// 		}
// 	}

// 	return -1
// }
