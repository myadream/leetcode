package leetcode_2

import (
	"fmt"
	"leetcode/src/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func outputListNode(nodes *ListNode) []int {
	list := make([]int, 0)
	for nodes != nil {
		list = append(list, nodes.Val)
		nodes = nodes.Next
	}

	fmt.Println("value: ", list)

	return list
}

func makeListNode(nodes []int) *ListNode {
	pre := &ListNode{
		Val:  0,
		Next: nil,
	}

	cur := pre

	for _, val := range nodes {
		node := &ListNode{
			Val:  val,
			Next: nil,
		}

		cur.Next = node
		cur = cur.Next
	}

	return pre.Next
}

func dataSet() []common.DataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var dataSets []common.DataSet[common.DataSetParamAssist, common.DataSetTarget]

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  makeListNode([]int{2, 4, 3}),
			Assist: makeListNode([]int{5, 6, 4}),
		},
		Target: common.DataSetTarget{
			Value: []int{7, 0, 8},
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  makeListNode([]int{0}),
			Assist: makeListNode([]int{0}),
		},
		Target: common.DataSetTarget{
			Value: []int{0},
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  makeListNode([]int{9, 9, 9, 9, 9, 9, 9}),
			Assist: makeListNode([]int{9, 9, 9, 9}),
		},
		Target: common.DataSetTarget{
			Value: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var handle []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]

	handle = append(handle, common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]{
		HandlerDataSet: func(dataSet common.DataSetParamAssist, target common.DataSetTarget) bool {

			res := outputListNode(addTwoNumbers(dataSet.Value.(*ListNode), dataSet.Assist.(*ListNode)))
			assert.Equal(
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

			return true
		},
	})

	return handle
}

func TestOne(Test *testing.T) {
	test := common.DataSetControl[common.DataSetParamAssist, common.DataSetTarget]{}
	test.Run(dataSet(), handle(Test))
}
