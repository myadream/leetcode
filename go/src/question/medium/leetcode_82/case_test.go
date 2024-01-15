package leetcode_82

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{1, 1, 1, 2, 3},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: []int{2, 3},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{1, 2, 3, 3, 4, 4, 5},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: []int{1, 2, 5},
		},
	})

	return dataSets
}

func makeListNode(nums []int) *ListNode {

	if len(nums) <= 0 {
		return nil
	}

	return &ListNode{
		Val:  nums[0],
		Next: makeListNode(nums[1:]),
	}
}

func outputListNode(head *ListNode) []int {
	var nums []int

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return nums
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := outputListNode(caseOne(makeListNode(data.SourceData.Value.([]int))))
		assert.Equal(
			t,
			res,
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)
	}
}
