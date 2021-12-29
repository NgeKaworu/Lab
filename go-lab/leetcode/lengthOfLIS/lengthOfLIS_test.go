package lengthoflis

import "testing"

func TestLengthOfLIS(t *testing.T) {
	testTable := []struct {
		nums []int
		res  int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{}, 0},
		{[]int{4, 10, 4, 3, 8, 9}, 3},
	}

	for _, tt := range testTable {
		actual := lengthOfLIS(tt.nums)
		if actual != tt.res {
			t.Errorf("lengthOfLIS(%v) = %v; expected %v", tt.nums, actual, tt.res)
		}
	}

}
