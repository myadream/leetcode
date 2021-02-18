package defangIPaddr

//https://leetcode-cn.com/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	var ans []byte

	for i := 0; i < len(address); i++ {
		if address[i] == 46 {
			ans = append(ans, 91)
			ans = append(ans, 46)
			ans = append(ans, 93)
		} else {
			ans = append(ans, address[i])
		}
	}

	return string(ans)
}
