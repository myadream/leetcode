package maxArea

import "testing"

func TestMaxArea(t *testing.T) {
	 val := []int{1,8,6,2,5,4,8,3,7}
	 target := 49

	 result := maxArea(val)
	 if result != target {
		 t.Fatalf("result != %d, result = %d", target, result)
	 }

	 //demo 2
	 val = []int{1,1}
	 target = 1

	 result = maxArea(val)
	 if result != target {
		 t.Fatalf("result != %d, result = %d", target, result)
	 }

	 //demo 3
	 val = []int{1,2,1}
	 target = 2

	 result = maxArea(val)
	 if result != target {
	 	t.Fatalf("result != %d, result = %d", target, result)
	 }

	 val = []int{2,3,4,5,18,17,6}
	 result = maxArea(val)
	 target = 17

	 if result != target {
	 	t.Fatalf("result != %d, result = %d", target, result)
	 }


}