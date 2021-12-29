package maxprofitk2

import "math"

func maxProfit(prices []int) int {
	dpI20, dpI10, dpI21, dpI11 := 0, 0, math.MinInt64, math.MinInt64
	for _, v := range prices {
		dpI20 = max(dpI20, dpI21+v)
		dpI21 = max(dpI21, dpI10-v)
		dpI10 = max(dpI10, dpI11+v)
		dpI11 = max(dpI11, -v)
	}
	return dpI20
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
