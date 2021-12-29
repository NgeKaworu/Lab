package reversepairs

func reversePairs(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	// split
	a1 := append([]int(nil), nums[:l/2]...)
	a2 := append([]int(nil), nums[l/2:]...)
	sol := reversePairs(a1) + reversePairs(a2)

	j := 0
	for _, v := range a1 {
		for j < len(a2) && v > 2*a2[j] {
			j++
		}
		sol += j
	}

	//merge
	p1, p2 := 0, 0
	for i := range nums {
		if p1 < len(a1) && (p2 == len(a2) || a1[p1] <= a2[p2]) {
			nums[i] = a1[p1]
			p1++
		} else {
			nums[i] = a2[p2]
			p2++
		}
	}

	return sol
}
