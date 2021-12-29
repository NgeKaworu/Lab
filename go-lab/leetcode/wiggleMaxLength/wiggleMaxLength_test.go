package wigglemaxlength

import "testing"

func TestWiggleMaxLength(t *testing.T) {
	tests := []struct {
		nums []int
		res  int
	}{

		{[]int{8, 7, 4, 9, 2, 5}, 5},
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		{[]int{1, 4, 7, 2, 5}, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{[]int{10, 10}, 1},
		{[]int{10, 10, 10}, 1},
		{[]int{3, 3, 3, 2, 5}, 3},
	}

	for _, v := range tests {
		acutal := wiggleMaxLength(v.nums)
		if acutal != v.res {
			t.Errorf("wiggleMaxLength(%v) = %v; expected %v", v.nums, acutal, v.res)
		}
	}
}
