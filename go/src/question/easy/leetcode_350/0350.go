package leetcode_350

func intersect(nums1 []int, nums2 []int) []int {
	var count = make(map[int]int32, len(nums1))
	var ans = []int{}

	for _, num := range nums1 {
		count[num] += 1
	}

	for _, num := range nums2 {
		if sum, ok := count[num]; ok && sum > 0 {
			count[num] -= 1
			ans = append(ans, num)
		}
	}

	return ans

}
