package leetcode_303

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {

			methods := dataSet.Value.([]string)
			datasets := dataSet.Assist.([][]int)
			targets := target.Value.([]int)

			var fn caseOne
			for index, method := range methods {
				if method == "NumArray" {
					fn = constructorCaseOne(datasets[index])
				} else {
					res := fn.SumRange(datasets[index][0], datasets[index][1])
					assert.Equal(
						t,
						res,
						targets[index],
						fmt.Sprintf(
							"case one: dataSet: %v, assist: %v, target: %v, res: %v",
							method,
							datasets[index],
							targets[index],
							res,
						),
					)
				}
			}

			return true
		},
	})

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {

			methods := dataSet.Value.([]string)
			datasets := dataSet.Assist.([][]int)
			targets := target.Value.([]int)

			var fn caseTwo
			for index, method := range methods {
				if method == "NumArray" {
					fn = constructorCaseTwo(datasets[index])
				} else {
					res := fn.SumRange(datasets[index][0], datasets[index][1])
					assert.Equal(
						t,
						res,
						targets[index],
						fmt.Sprintf(
							"case one: dataSet: %v, assist: %v, target: %v, res: %v",
							method,
							datasets[index],
							targets[index],
							res,
						),
					)
				}
			}

			return true
		},
	})

	return handle
}
