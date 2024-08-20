package perfectsquarecheck

import "math"

func solution(n int64) int64 {
	sqrt := int64(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		return (sqrt + 1) * (sqrt + 1)
	}

	return -1
}
