package leetcode_82

type ListNode struct {
	Val  int
	Next *ListNode
}

func caseOne(head *ListNode) *ListNode {
	var tail = &ListNode{Next: head}
	var pre, cur = tail, head

	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}

		if pre.Next == cur {
			pre = cur
		} else {
			pre.Next = cur.Next
		}

		cur = cur.Next
	}

	return tail.Next
}
