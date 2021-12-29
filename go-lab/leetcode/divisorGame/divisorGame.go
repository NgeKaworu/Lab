// 1025. 除数博弈
// https://leetcode-cn.com/problems/divisor-game/

package divisorgame

// 递推版
func divisorGame(N int) bool {

	dp := make([]bool, N+1)
	dp[0], dp[1] = false, false

	for i := 2; i <= N; i++ {
		divs := getAllDivisor(i)
		canWin := false
		for _, v := range divs {
			if !dp[i-v] {
				canWin = true
				break
			}
		}
		dp[i] = canWin
	}
	return dp[N]
}

// 数学版
// func divisorGame(N int) bool {
// 	return N%2 == 0
// }

func getAllDivisor(n int) []int {
	res := []int{1}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			res = append(res, i)
		}
	}
	return res
}
