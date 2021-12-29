// 118. 杨辉三角
// https://leetcode-cn.com/problems/pascals-triangle/

package generate

func generate(numRows int) [][]int {
	sol := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		sol[i] = make([]int, i+1)
		for k := range sol[i] {
			if k == 0 || k == i {
				sol[i][k] = 1
			} else {
				sol[i][k] = sol[i-1][k] + sol[i-1][k-1]
			}

		}
	}
	return sol
}
