package maximum

import "fmt"

//https://leetcode-cn.com/problems/maximum-lcci/
func maximum(a int, b int) int {
	fmt.Printf("%b, %b, %b", a, b, a & b)
	return a & b
}
