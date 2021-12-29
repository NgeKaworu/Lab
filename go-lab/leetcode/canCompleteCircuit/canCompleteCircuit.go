package cancompletecircuit

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for i := 0; i < l; i++ {
		if gas[i] < cost[i] {
			continue
		}
		g := 0
		for j := 0; j < l; j++ {
			cur := (i + j) % l
			g = g + gas[cur] - cost[cur]
			if g < 0 {
				// 剪枝
				if cur-1 > i {
					i = cur
				}
				break
			}
		}
		if g >= 0 {
			return i
		}
	}
	return -1
}
