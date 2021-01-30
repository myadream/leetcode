package SubrectangleQueries

import "testing"

func TestSubrectangleQueries(t *testing.T) {
	rectangle := [][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}

	subrectangleQueries := Constructor(rectangle)

	var val = subrectangleQueries.GetValue(0, 2)
	if val != 1 {
		t.Fatal("getValue not eq 1")
	}

	subrectangleQueries.UpdateSubrectangle(0, 0, 3, 2, 5)
	val = subrectangleQueries.GetValue(0, 2)
	if val != 5 {
		t.Fatal("getValue not eq 5")
	}

	val = subrectangleQueries.GetValue(3, 2)
	if val != 5 {
		t.Fatal("getValue not eq 5")
	}

	subrectangleQueries.UpdateSubrectangle(3, 0, 3, 2, 10)
	val = subrectangleQueries.GetValue(3, 1)
	if val != 10 {
		t.Fatal("getValue not eq 10")
	}

	val = subrectangleQueries.GetValue(0, 2)
	if val != 5 {
		t.Fatal("getValue not eq 5")
	}

}