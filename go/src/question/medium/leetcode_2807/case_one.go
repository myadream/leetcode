package leetcode_2807

func CaseOne(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	nodes := head

	for nodes.Next != nil {
		ans := gcd(nodes.Val, nodes.Next.Val)
		node := &ListNode{Val: ans, Next: nodes.Next}
		nodes.Next = node
		nodes = node.Next
	}

	return head
}

func CaseTwo(head *ListNode) *ListNode {
	if head.Next != nil {
		CaseTwo(head.Next)
		head.Next = &ListNode{Val: gcd(head.Val, head.Next.Val), Next: head.Next}
	}

	return head
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}
