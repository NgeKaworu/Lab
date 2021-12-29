// https://leetcode-cn.com/problems/count-primes/
// 204. 计数质数
// https://blog.csdn.net/afei__/article/details/80638460
// 快速判断质数

package countprimes

import "math"

func countPrimes(n int) int {
	var sol int
	for n > 1 {
		n--
		if isPrime(n) {
			sol++
		}
	}
	return sol
}

func isPrime(num int) bool {
	if num <= 3 {
		return num > 1
	}
	// 不在6的倍数两侧的一定不是质数
	if num%6 != 1 && num%6 != 5 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 5; i <= sqrt; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

// func countPrimes(n int) int {
// 	//厄拉多塞筛法
// 	count := 0
// 	signs := make([]bool, n)
// 	for i := 2; i < n; i++ {
// 		if signs[i] {
// 			continue
// 		}
// 		count++
// 		for j := i + i; j < n; j += i {
// 			signs[j] = true
// 		}
// 	}
// 	return count
// }
