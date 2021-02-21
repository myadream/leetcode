package findShortestSubArray

type PII struct {
	x, y int
}

//https://leetcode-cn.com/problems/degree-of-an-array/
func findShortestSubArray(nums []int) (ans int) {
	var mp = make(map[int]*PII)
	var ma int

	for i := 0; i < len(nums); i++ {
		if mp[nums[i]] == nil {
			mp[nums[i]] = &PII{
				x: 1,
				y: i,
			}
		} else {
			 mp[nums[i]].x++
		}

		if ma < mp[nums[i]].x {
			ma = mp[nums[i]].x
			ans = i - mp[nums[i]].y + 1
		}

		if ma == mp[nums[i]].x {
			ans = min(ans, i - mp[nums[i]].y + 1)
		}
	}
	return
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}