package defangIPaddr

import "bytes"

//https://leetcode-cn.com/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	var ans bytes.Buffer

	for i := 0; i < len(address); i++ {
		if address[i] == 46 {
			ans.WriteByte(91)
			ans.WriteByte(46)
			ans.WriteByte(93)
		} else {
			ans.WriteByte(address[i])
		}
	}

	return ans.String()
}
