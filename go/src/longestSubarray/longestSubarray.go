package longestSubarray


/**
func longestSubarray(nums []int, limit int) (ans int) {
    var minQ, maxQ []int
    left := 0
    for right, v := range nums {
        for len(minQ) > 0 && minQ[len(minQ)-1] > v {
            minQ = minQ[:len(minQ)-1]
        }
        minQ = append(minQ, v)
        for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
            maxQ = maxQ[:len(maxQ)-1]
        }
        maxQ = append(maxQ, v)
        for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
            if nums[left] == minQ[0] {
                minQ = minQ[1:]
            }
            if nums[left] == maxQ[0] {
                maxQ = maxQ[1:]
            }
            left++
        }
        ans = max(ans, right-left+1)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/solution/jue-dui-chai-bu-chao-guo-xian-zhi-de-zui-5bki/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */

func longestSubarray(nums []int, limit int) (ans int) {
	var minQ, maxQ []int
	var left = 0

	for right, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)

		for len(maxQ) > 0 && maxQ[len(maxQ) - 1] < v {
			maxQ = maxQ[:len(maxQ) - 1]
		}

		maxQ = append(maxQ, v)
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		ans = max(ans, right-left+1)
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
