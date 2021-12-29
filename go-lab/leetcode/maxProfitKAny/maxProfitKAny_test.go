package maxprofitk2

import "testing"

// TestMaxProfit test
func TestMaxProfit(t *testing.T) {
	testTable := []struct {
		k      int
		prices []int
		res    int
	}{
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
		{1, []int{1, 2}, 1},
	}

	for _, tt := range testTable {
		actual := maxProfit(tt.k, tt.prices)
		if actual != tt.res {
			t.Errorf("removeKDigits(%v, %v) = %v; expected %v", tt.k, tt.prices, actual, tt.res)
		}
	}
}
