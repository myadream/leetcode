package addToArrayForm

//https://leetcode-cn.com/problems/add-to-array-form-of-integer/solution/shu-zu-xing-shi-de-zheng-shu-jia-fa-by-l-jljp/
func addToArrayForm(A []int, K int) []int {
	var ans []int
	for i := len(A) - 1; i >= 0 || K > 0; i-- {
		if i >= 0 {
			K += A[i]
		}

		ans = append(ans, K % 10)
		K /= 10
	}

	return reverse(ans)
}

func reverse(A []int) []int {
	for i, n := 0, len(A); i < n/2; i++ {
		A[i], A[n-1-i] = A[n-1-i], A[i]
	}

	return A
}

