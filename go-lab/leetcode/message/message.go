package message

// 暴力递归 T(n) = O(n * n!) S(n) = O(1)
// func massage(nums []int) int {
// 	if len(nums) <= 2 {
// 		return max(nums...)
// 	}
// 	res := 0
// 	for i := 0; i < len(nums)-1; i++ {
// 		res = max(res, nums[i]+massage(nums[i+2:]))
// 	}
// 	return res
// }

// dp T(n) = O(n) S(n) = O(n)
// func massage(nums []int) int {
// 	l := len(nums)
// 	if l <= 2 {
// 		return max(nums...)
// 	}
// 	dp := make([]int, l)
// 	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
// 	for i := 2; i < l; i++ {
// 		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
// 	}
// 	return dp[l-1]
// }

// dp S(n) = O(1) T(n) = O(n)
func massage(nums []int) int {
	pre, res := 0, 0
	for i := 0; i < len(nums); i++ {
		res, pre = max(res, pre+nums[i]), res
	}
	return res
}

func max(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	res := n[0]
	for _, v := range n {
		if v > res {
			res = v
		}
	}
	return res
}
