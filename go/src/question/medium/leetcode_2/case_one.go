package leetcode_2

// CaseOne /**
func CaseOne(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}

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

		carry = sum / 10
		sum %= 10
		cur.Next = &ListNode{
			Val: sum,
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
