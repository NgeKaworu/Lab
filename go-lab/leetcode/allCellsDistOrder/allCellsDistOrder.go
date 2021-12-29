package allcellsdistorder

import "sort"

func allCellsDistOrder(n, m, r0, c0 int) [][]int {
	ans := make([][]int, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = append(ans, []int{i, j})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		a, b := ans[i], ans[j]
		return abs(a[0]-r0)+abs(a[1]-c0) < abs(b[0]-r0)+abs(b[1]-c0)
	})
	return ans
}

// 桶排序
// func allCellsDistOrder(n, m, r0, c0 int) [][]int {
//     maxDist := max(r0, n-1-r0) + max(c0, m-1-c0)
//     buckets := make([][][]int, maxDist+1)

//     for i := 0; i < n; i++ {
//         for j := 0; j < m; j++ {
//             dist := abs(i-r0) + abs(j-c0)
//             buckets[dist] = append(buckets[dist], []int{i, j})
//         }
//     }

//     ans := make([][]int, 0, n*m)
//     for _, bucket := range buckets {
//         ans = append(ans, bucket...)
//     }
//     return ans
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
