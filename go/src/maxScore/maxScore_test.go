package maxScore

import "testing"

func TestMaxScore(t *testing.T) {
	var cardPoints = []int{1,2,3,4,5,6,1}
	var k = 3
	var target = 12

	var result = maxScore(cardPoints, k)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	cardPoints = []int{2, 2, 2, 2}
	k = 2
	target = 4

	result = maxScore(cardPoints, k)
	if target != result {
		t.Fatalf("result: %d", result)
	}


	cardPoints = []int{9,7,7,9,7,7,9}
	k = 7
	target = 55

	result = maxScore(cardPoints, k)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	cardPoints = []int{1, 1000, 1}
	k = 1
	target = 1

	result = maxScore(cardPoints, k)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	cardPoints = []int{1,79,80,1,1,1,200,1}
	k = 3
	target = 202

	result = maxScore(cardPoints, k)
	if target != result {
		t.Fatalf("result: %d", result)
	}

}