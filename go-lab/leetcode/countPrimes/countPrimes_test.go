package countprimes

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		n   int
		sol int
	}{
		{10, 4},
		{0, 0},
		{1, 0},
		{2, 0},
	}
	for _, tt := range tests {
		actual := countPrimes(tt.n)
		if actual != tt.sol {
			t.Errorf("findMinArrowShots(%v) = %v, expected %v", tt.n, actual, tt.sol)
		}
	}
}
