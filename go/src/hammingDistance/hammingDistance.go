package hammingDistance

import "fmt"

//https://leetcode-cn.com/problems/hamming-distance/
func hammingDistance(x int, y int) (ans int) {
	var xor = x ^ y

	for xor > 0 {
		fmt.Println(xor, xor - 1, xor & xor - 1)

		ans++
		xor &= xor - 1
	}

	return
}
