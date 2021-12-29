package message

import "testing"

func TestMassage(t *testing.T) {
	tests := []struct {
		nums []int
		res  int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 4, 5, 3, 1, 1, 3}, 12},
		{[]int{}, 0},
		{[]int{1, 3, 1}, 3},
	}

	for _, v := range tests {
		actual := massage(v.nums)
		if v.res != actual {
			t.Errorf("massage(%v) = %v; expected %v", v.nums, actual, v.res)
		}
	}
}
