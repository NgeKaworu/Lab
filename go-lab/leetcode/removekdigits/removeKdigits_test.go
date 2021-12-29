package removekdigits

import "testing"

// TestRemoveKDigits test func
func TestRemoveKDigits(t *testing.T) {
	testTable := []struct {
		num string
		k   int
		res string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 1, "0"},
		{"112", 1, "11"},
	}

	for _, tt := range testTable {
		actual := removeKdigits(tt.num, tt.k)
		if actual != tt.res {
			t.Errorf("removeKDigits(%v, %d) = %v; expected %v", tt.num, tt.k, actual, tt.res)
		}
	}
}
