package intersection

//https://leetcode-cn.com/problems/intersection-of-two-arrays/
func intersection(nums1 []int, nums2 []int) []int {
	var ans = map[int]int{}
	var k = len(nums1)

	for _, v := range nums1 {
		ans[v] = 1
	}

	for _, v := range nums2 {
		if ans[v] == 1 {
			ans[v] = 0
			nums1 = append(nums1, v)
		}
	}

	return nums1[k:]
}