package getDecimalValue

type ListNode struct {
	Val int
	Next *ListNode
}

//https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/
func getDecimalValue(head *ListNode) (ans int) {
	for head != nil {
		ans = (ans << 1) | head.Val
		head = head.Next
	}

	return
}