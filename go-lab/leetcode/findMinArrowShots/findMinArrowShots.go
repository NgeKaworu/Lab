// 452. 用最少数量的箭引爆气球
// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/

package findminarrowshots

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) < 1 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res := 1
	arrow := points[0][1]
	for _, v := range points {
		if v[0] > arrow {
			arrow = v[1]
			res++
		}

	}
	return res
}
