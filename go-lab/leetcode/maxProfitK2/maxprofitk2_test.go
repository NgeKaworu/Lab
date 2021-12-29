package maxprofitk2

import "testing"

// TestMaxProfit test
func TestMaxProfit(t *testing.T) {
	testTable := []struct {
		prices []int
		res    int
	}{
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
	}

	for _, tt := range testTable {
		actual := maxProfit(tt.prices)
		if actual != tt.res {
			t.Errorf("removeKDigits(%v) = %v; expected %v", tt.prices, actual, tt.res)
		}
	}
}
