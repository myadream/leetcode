package getDecimalValue

import (
	"testing"
)

func TestGetDecimalValue(t *testing.T) {
	var temp = []int{1, 0, 1}
	var target = 5
	var result = getDecimalValue(getListNode(temp))

	if target != result {
		t.Fatalf("result: %v", result)
	}

	temp = []int{0}
	target = 0
	result = getDecimalValue(getListNode(temp))

	if target != result {
		t.Fatalf("result: %v", result)
	}

	temp = []int{1}
	target = 1
	result = getDecimalValue(getListNode(temp))

	if target != result {
		t.Fatalf("result: %v", result)
	}


	temp = []int{0, 0}
	target = 0
	result = getDecimalValue(getListNode(temp))

	if target != result {
		t.Fatalf("result: %v", result)
	}

	temp = []int{1,0,0,1,0,0,1,1,1,0,0,0,0,0,0}
	target = 18880
	result = getDecimalValue(getListNode(temp))

	if target != result {
		t.Fatalf("result: %v", result)
	}
}

func getListNode(temp []int) *ListNode {
	var head = &ListNode{}
	for _, v := range temp {
		node := &ListNode{
			Val: v,
			Next: nil,
		}

		node.Next = head
		head = node
	}

	return head
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}