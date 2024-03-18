package leetcode_303

type caseTwo struct {
	nums   []int
	presum []int
}

func constructorCaseTwo(nums []int) caseTwo {

	presum := make([]int, len(nums)+1)

	presum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		presum[i] = presum[i-1] + nums[i]
	}

	return caseTwo{
		nums:   nums,
		presum: presum,
	}
}

func (this *caseTwo) SumRange(left int, right int) int {
	return this.presum[right] - this.presum[left] + this.nums[left]
}
