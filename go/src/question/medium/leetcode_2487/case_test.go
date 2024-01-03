package template

import (
	"fmt"
	"leetcode/src/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeNodes(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}

	node := &ListNode{Val: nums[0]}
	node.Next = makeNodes(nums[1:])

	return node
}

func outputNodes(node *ListNode) []int {
	if node == nil {
		return []int{}
	}

	var nums []int
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}

	return nums
}

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value: []int{5, 2, 13, 3, 8},
		}, TargetData: common.TDefault{
			Value: []int{13, 8},
		},
	})

	return dataSets
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		val := makeNodes(data.SourceData.Value.([]int))
		res := outputNodes(CaseOne(val))
		assert.Equal(
			t,
			res,
			data.TargetData.Value.([]int),
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				outputNodes(val),
				data.TargetData.Value,
				res,
			),
		)
	}
}

func TestCaseTwo(t *testing.T) {
	for _, data := range dataSet() {
		val := makeNodes(data.SourceData.Value.([]int))
		res := outputNodes(CaseTwo(val))
		assert.Equal(
			t,
			res,
			data.TargetData.Value.([]int),
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				outputNodes(val),
				data.TargetData.Value,
				res,
			),
		)
	}
}
