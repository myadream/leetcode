package leetcode_303

type caseOne struct {
	nums []int
	n    int
}

func constructorCaseOne(nums []int) caseOne {
	return caseOne{
		nums: nums,
		n:    len(nums),
	}
}

func (this *caseOne) SumRange(left int, right int) int {
	ans := 0
	for left <= right && left <= this.n {
		ans += this.nums[left]
		left += 1
	}

	return ans
}
