package maximumWealth

import "testing"

func TestMaximumWealth(t *testing.T) {
	var accounts = [][]int{{1, 2, 3}, {3, 2, 1}}
	var target = 6

	var result = maximumWealth(accounts)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	accounts = [][]int{{1, 5}, {7, 3}, {3, 5}}
	target = 10

	result = maximumWealth(accounts)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	accounts = [][]int{{2, 7, 8}, {7, 1, 3}, {1, 9, 5}}
	target = 17

	result = maximumWealth(accounts)
	if result != target {
		t.Fatalf("result: %d", result)
	}
}