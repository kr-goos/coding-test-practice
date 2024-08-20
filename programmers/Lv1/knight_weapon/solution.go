package knightweapon

import "math"

// 개선 전
func solution(number int, limit int, power int) int {
	var total int
	for i := 1; i <= number; i++ {
		var weight int
		for j := 1; j <= i; j++ {
			if i%j != 0 {
				continue
			}
			weight++

			if weight > limit {
				weight = power
				break
			}
		}
		total += weight
	}
	return total
}

// 개선 후
func solution2(number int, limit int, power int) int {
	var total int
	for i := 1; i <= number; i++ {
		var weight int
		sqrtI := int(math.Sqrt(float64(i)))
		for j := 1; j <= sqrtI; j++ {
			if i%j != 0 {
				continue
			}

			if j*j == i {
				weight++
			} else {
				weight += 2
			}

			if weight > limit {
				weight = power
				break
			}
		}
		total += weight
	}
	return total
}
