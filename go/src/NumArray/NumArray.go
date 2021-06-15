package NumArray

type NumArray struct {
	nums []int
}


func Constructor(nums []int) NumArray {
	n := NumArray{
		nums: make([]int, len(nums) + 1),
	}

	for k, v := range nums {
		n.nums[k + 1] = n.nums[k] + v
	}

	return n
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.nums[j + 1] - this.nums[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

/**

type NumArray struct {
	nums []int
}


func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}


func (this *NumArray) SumRange(i int, j int) (ans int) {
	for ;i <= j; i++ {
		ans += this.nums[i]
	}

	return
}

 */