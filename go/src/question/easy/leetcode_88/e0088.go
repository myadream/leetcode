package leetcode_88

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {

	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}

	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i := len(nums1) - 1
	m--
	n--

	for n >= 0 {
		if m >= 0 && nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}

		i--
	}
}
