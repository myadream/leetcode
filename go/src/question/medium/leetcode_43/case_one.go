package leetcode_67

import "math/big"

func CaseOne(a string, b string) string {

	ai, _ := new(big.Int).SetString(a, 10)
	bi, _ := new(big.Int).SetString(b, 10)

	return ai.Mul(ai, bi).Text(10)
}
