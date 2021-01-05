package TwoNum

import (
	"reflect"
	"testing"
)

func TestTwoNum(t *testing.T) {
	 var nums []int = []int{2,7,11,15}
	 var target int = 9

	 var result []int = twoSum(nums, target)
	 var successResult []int = []int{0, 1}

	 if !reflect.DeepEqual(result, successResult) {
	 	t.Fatal("匹配错误")
	 }
}