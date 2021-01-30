package decompressRLElist


//https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
func decompressRLElist(nums []int) []int {
	var ans []int

	for i := 0; i < len(nums); i += 2 {
		t := make([]int, nums[i])
		for k, _ := range t {
			t[k] = nums[i + 1]
		}

		ans = append(ans, t...)
	}

	return ans
}