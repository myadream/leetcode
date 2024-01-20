package leetcode_128

func caseOne(nums []int) int {

	l := len(nums)
	numSet := make(map[int]bool, l)
	longestStreak := 0
	for i := l - 1; i >= 0; i-- {
		numSet[nums[i]] = true
	}

	for num := range numSet {
		if !numSet[num-1] {
			currentStreak := 1
			for numSet[num+1] {
				num++
				currentStreak++
			}

			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
