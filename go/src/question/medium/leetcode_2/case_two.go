package leetcode_2

func CaseTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	var fake ListNode
	var head *ListNode = &fake
	var overflow int
	for l1 != nil && l2 != nil {
		s := l1.Val + l2.Val + overflow

		overflow = 0
		if s >= 10 {
			overflow = 1
		}
		head.Next = &ListNode{
			Val: s % 10,
		}
		head = head.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	var l *ListNode
	if l1 != nil {
		l = l1
	} else {
		l = l2
	}

	for l != nil {
		s := l.Val + overflow
		if s >= 10 {
			overflow = 1
		} else {
			overflow = 0
		}
		head.Next = &ListNode{
			Val: s % 10,
		}
		head = head.Next
		l = l.Next
	}

	if overflow != 0 {
		head.Next = &ListNode{
			Val: overflow,
		}
	}

	return fake.Next
}
