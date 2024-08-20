package sumofsignedintegers

func solution(absolutes []int, signs []bool) int {

	var total int
	for i, v := range absolutes {
		if !signs[i] {
			v *= -1
		}
		total += v
	}

	return total
}
