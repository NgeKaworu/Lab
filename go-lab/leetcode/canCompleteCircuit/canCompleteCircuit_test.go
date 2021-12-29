package cancompletecircuit

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas  []int
		cost []int
		res  int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
	}

	for _, tt := range tests {
		auctal := canCompleteCircuit(tt.gas, tt.cost)
		if auctal != tt.res {
			t.Errorf("canCompleteCircuit(%v, %v) = %v, expected %v", tt.gas, tt.cost, auctal, tt.res)
		}
	}
}
