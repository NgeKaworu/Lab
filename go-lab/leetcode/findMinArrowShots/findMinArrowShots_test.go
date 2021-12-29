package findminarrowshots

import (
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		points [][]int
		res    int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{[][]int{{1, 2}}, 1},
		{[][]int{{2, 3}, {2, 3}}, 1},
	}
	for _, tt := range tests {
		actual := findMinArrowShots(tt.points)
		if actual != tt.res {
			t.Errorf("findMinArrowShots(%v) = %v, expected %v", tt.points, actual, tt.res)
		}
	}
}
