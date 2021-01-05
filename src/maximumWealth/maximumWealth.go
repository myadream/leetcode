package maximumWealth

//https://leetcode-cn.com/problems/richest-customer-wealth/
func maximumWealth(accounts [][]int) int {
	for ak, av := range accounts {
		for i := 1; i < len(av); i++ {
			accounts[ak][0] += accounts[ak][i]
		}

		if accounts[ak][0] > accounts[0][0] {
			accounts[0][0] = accounts[ak][0]
		}
	}


	return accounts[0][0]
}
