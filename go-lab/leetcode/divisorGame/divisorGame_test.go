package divisorgame

import "testing"

func TestDivisorGame(t *testing.T) {
	tests := []struct {
		N   int
		res bool
	}{
		{2, true},
		{3, false},
		{4, true},
	}

	for _, tt := range tests {
		actual := divisorGame(tt.N)
		if actual != tt.res {
			t.Errorf("divsorGame(%v) = %v, expected %v", tt.N, actual, tt.res)
		}
	}
}
