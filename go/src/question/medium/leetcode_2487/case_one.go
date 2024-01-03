package template

type ListNode struct {
	Val  int
	Next *ListNode
}

func CaseOne(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = CaseOne(head.Next)
	if head.Next != nil && head.Val < head.Next.Val {
		head = head.Next
	}

	return head
}

func reverse(node *ListNode) *ListNode {
	var prev, curr *ListNode
	curr = node
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func CaseTwo(head *ListNode) *ListNode {
	head = reverse(head)

	var max, curr *ListNode = head, head

	for curr.Next != nil {
		curr = curr.Next
		if curr.Val >= max.Val {
			max.Next = curr
			max = curr
		}
	}

	max.Next = nil

	return reverse(head)
}
