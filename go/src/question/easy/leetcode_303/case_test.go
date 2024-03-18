package leetcode_303

import (
	"leetcode/src/common"
	"testing"
)

func TestCaseOne(t *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(t), dataSet())
}
