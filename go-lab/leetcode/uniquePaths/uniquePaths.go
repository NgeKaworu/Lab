package uniquepaths

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	// init
	for k := range dp {
		dp[k] = make([]int, n)
		if k == 0 {
			for k2 := range dp[k] {
				dp[k][k2] = 1
			}
		}
		dp[k][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
