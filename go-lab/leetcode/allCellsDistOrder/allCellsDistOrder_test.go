package allcellsdistorder

import (
	"fmt"
	"testing"
)

func TestAllCellsDistOrder(t *testing.T) {
	testTable := []struct {
		R   int
		C   int
		r0  int
		c0  int
		res [][]int
	}{
		{2, 2, 0, 1, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
		{2, 3, 1, 2, [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}},
	}

	for _, tt := range testTable {
		actual := allCellsDistOrder(tt.R, tt.C, tt.r0, tt.c0)
		fmt.Printf("lengthOfLIS(%v, %v, %v, %v) = %v; expected %v \n", tt.R, tt.C, tt.r0, tt.c0, actual, tt.res)
	}

}
