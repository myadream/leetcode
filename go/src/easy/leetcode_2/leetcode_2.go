package leetcode_2

/**
 * Definition for singly-linked list.
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{
		Next: nil,
		Val:  0,
	}

	cur := pre

	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry /= 10
		sum %= 10
		cur.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		cur = cur.Next
	}

	if carry == 1 {
		cur.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return pre.Next
}
