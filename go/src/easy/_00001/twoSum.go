package _00001

func twoSum(nums []int, target int) []int {
	var res = make(map[int]int, len(nums)>>1)
	for k, v := range nums {
		if res[v] > 0 {
			return []int{res[v] - 1, k}
		} else {
			res[target-v] = k + 1
		}
	}

	return nil
}
