package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n, res int
	}{
		{3, 2, 3},
		{7, 3, 28},
	}

	for _, v := range tests {
		actual := uniquePaths(v.m, v.n)
		if actual != v.res {
			t.Errorf("uniquePaths(%v, %v) = %v; expected %v ", v.m, v.n, actual, v.res)
		}
	}
}
