package leetcode_2

func CaseThree(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	tail := head

	for l1 != nil || l2 != nil {
		if l1 != nil {
			tail.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tail.Val += l2.Val
			l2 = l2.Next
		}

		tail = addNode(tail, l1 != nil || l2 != nil)
	}

	var l *ListNode = l1
	if l1 == nil {
		l = l2
	}

	for l != nil {
		tail.Val += l.Val
		l = l.Next

		tail = addNode(tail, l != nil)
	}

	return head
}

func addNode(tail *ListNode, state bool) *ListNode {
	beyond := 0
	if tail.Val >= 10 {
		tail.Val %= 10
		beyond = 1
	}

	if state || beyond > 0 {
		tail.Next = &ListNode{
			Val: beyond,
		}

		tail = tail.Next
	}

	return tail
}
