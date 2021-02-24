package containsDuplicate

import "sort"

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i - 1] {
			return true
		}
	}

	return false
}

//哈希表
//func containsDuplicate(nums []int) bool {
//	var m = make(map[int]int, len(nums))
//
//	for _, v := range nums {
//		m[v]++
//
//		if m[v] >= 2 {
//			return true
//		}
//	}
//
//	return false
//}