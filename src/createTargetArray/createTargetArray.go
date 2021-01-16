package createTargetArray

type Node struct {
	val int
	next *Node
}

func insert(head *Node, idx, val int) {
	p := head

	for i := 0; i < idx; i++ {
		p = p.next
	}

	var t = p.next
	p.next = &Node{
		val,
		nil,
	}

	p.next.next = t
}

func createTargetArray(nums []int, index []int) []int {
	i := 0
	n := len(index)
	node := &Node{}


	for i = 0; i < n; i++ {
		insert(node, index[i], nums[i])
	}

	i = 0
	res := make([]int, n)
	p := node.next

	for p != nil {
		res[i] = p.val
		i++
		p = p.next
	}

	return res
}