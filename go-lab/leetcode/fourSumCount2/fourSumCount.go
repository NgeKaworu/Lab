package foursumcount2

// 四数相加 II
// 题解 https://leetcode-cn.com/problems/4sum-ii/solution/si-shu-xiang-jia-ii-by-leetcode-solution/

// 回溯又双叒叕超时了
// func fourSumCount(A []int, B []int, C []int, D []int) int {
// 	var res int
// 	var backtrack func(dict [][]int, row, tmp int)
// 	backtrack = func(dict [][]int, row, tmp int) {
// 		if row == len(dict) {
// 			if tmp == 0 {
// 				res++
// 			}
// 			return
// 		}

// 		for _, v := range dict[row] {
// 			tmp += v
// 			backtrack(dict, row+1, tmp)
// 			tmp -= v
// 		}
// 	}

// 	d := [][]int{A, B, C, D}
// 	backtrack(d, 0, 0)
// 	return res
// }

// 哈希表 不明觉厉
func fourSumCount(a, b, c, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w]
		}
	}
	return
}
