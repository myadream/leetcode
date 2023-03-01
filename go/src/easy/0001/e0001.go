package _0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)>>1)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return []int{m[nums[i]], i}
		} else {
			m[target-nums[i]] = i
		}
	}

	return nil
}
