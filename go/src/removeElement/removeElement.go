package removeElement

import "fmt"

//https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	var res int

	for i, _ := range nums {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}

	fmt.Println(nums)

	return res
}