package monotoneincreasingdigits

import "testing"

func TestMonotoneIncreasingDigits(t *testing.T) {
	tests := []struct {
		N, sol int
	}{
		{10, 9},
		{1234, 1234},
		{332, 299},
		{12341234, 12339999},
	}

	for _, v := range tests {
		actual := monotoneIncreasingDigits(v.N)
		if actual != v.sol {
			t.Errorf("monotoneIncreasingDigits(%v) = %v; expected %v", v.N, actual, v.sol)
		}

	}
}
