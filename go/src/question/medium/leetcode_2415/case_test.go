package leetcode_2415

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func outputTreeNode(nodes *TreeNode) []int {
	list := make([]int, 0)
	for nodes != nil {
		list = append(list, nodes.Val)
		nodes = nodes.Next
	}

	fmt.Println("value: ", list)

	return list
}

func recursion(cur *TreeNode, nodes []int, index int) *TreeNode {
	if len(nodes) < index {
		return nil
	}

	l := &TreeNode{
		Val: nodes[index],
	}

	r := &TreeNode{
		Val: nodes[index+1],
	}

	cur.Left = recursion(l, nodes, index)
	cur.Right = recursion(r, nodes, index)

	return cur
}

func makeTreeNode(nodes []int) *TreeNode {
	pre := &TreeNode{
		Val: nodes[1],
	}

	cur := pre
	cur = recursion(cur, nodes, 1)

	return pre
}

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  makeListNode([]int{0}),
			Assist: makeListNode([]int{7, 3}),
		}, TargetData: common.TDefault{
			Value: []int{7, 3},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  makeListNode([]int{2, 4, 3}),
			Assist: makeListNode([]int{5, 6, 4}),
		}, TargetData: common.TDefault{
			Value: []int{7, 0, 8},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  makeListNode([]int{0}),
			Assist: makeListNode([]int{0}),
		}, TargetData: common.TDefault{
			Value: []int{0},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  makeListNode([]int{9, 9, 9, 9, 9, 9, 9}),
			Assist: makeListNode([]int{9, 9, 9, 9}),
		}, TargetData: common.TDefault{
			Value: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			res := outputListNode(CaseOne(dataSet.Value.(*ListNode), dataSet.Assist.(*ListNode)))
			return assert.Equal(
				t,
				res,
				target.Value.([]int),
				fmt.Sprintf(
					"case one: dataSet: %v, assist: %v, target: %v, res: %v",
					outputListNode(dataSet.Value.(*ListNode)),
					outputListNode(dataSet.Assist.(*ListNode)),
					target,
					res,
				),
			)
		},
	})

	return handle
}

func TestLeetCode2415(Test *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(Test), dataSet())
}
