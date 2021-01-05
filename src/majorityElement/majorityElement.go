package majorityElement

func majorityElement(nums []int) int {
	var major,cnt int

	for i := range nums {
		if cnt == 0 {
			major = nums[i]
			cnt++
		} else if major == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}

	if cnt > 0 {
		t := 0
		for i := range nums {
			if major == nums[i] {
				t++
			}
		}

		if t > (len(nums) / 2) {
			return major
		}
	}

	return -1
}
