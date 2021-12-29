package maxprofitk2

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var res int
	//第一维度表示日期，第二维度表示当前交易了k次，第三维度表示持仓状态，0-空仓，1-有票
	profit := make([][][2]int, len(prices), len(prices))
	for i := range profit {
		profit[i] = make([][2]int, k+1, k+1)
		for j := range profit[i] {
			profit[i][j] = [2]int{0, -prices[0]}
		}
	}

	for i := 1; i < len(profit); i++ {
		profit[i][0][0] = profit[i-1][0][0]
		profit[i][0][1] = max(profit[i-1][0][1], profit[i-1][0][0]-prices[i])
		for j := 1; j <= k; j++ {
			profit[i][j][0] = max(profit[i-1][j][0], profit[i-1][j-1][1]+prices[i])
			profit[i][j][1] = max(profit[i-1][j][1], profit[i-1][j][0]-prices[i])
		}
		res = max(res, profit[i][k][0])
	}

	return res
}

func max(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}
