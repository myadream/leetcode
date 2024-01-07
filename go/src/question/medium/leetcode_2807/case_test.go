package leetcode_2807

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func makeListNode(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}

	node := &ListNode{Val: nums[0]}
	node.Next = makeListNode(nums[1:])

	return node
}

func outputListNode(nodes *ListNode) []int {
	nums := []int{}

	for nodes != nil {
		nums = append(nums, nodes.Val)
		nodes = nodes.Next
	}

	return nums
}

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{18, 6, 10, 3},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: []int{18, 6, 6, 2, 10, 1, 3},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{7},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: []int{7},
		},
	})

	return dataSets
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := CaseOne(makeListNode(data.SourceData.Value.([]int)))
		assert.Equal(
			t,
			outputListNode(res),
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				outputListNode(res),
			),
		)
	}
}

func TestCaseTow(t *testing.T) {
	for _, data := range dataSet() {
		res := CaseTwo(makeListNode(data.SourceData.Value.([]int)))
		assert.Equal(
			t,
			outputListNode(res),
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				outputListNode(res),
			),
		)
	}
}
