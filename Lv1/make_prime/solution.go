package makeprime

import "math"

func solution(nums []int) int {

	var primeCount int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				v := nums[i] + nums[j] + nums[k]
				isPrime := true
				for idx := 2; idx <= int(math.Sqrt(float64(v))); idx++ {
					if v%idx == 0 {
						isPrime = false
						break
					}
				}
				if isPrime {
					primeCount++
				}
			}
		}
	}

	return primeCount
}
