package dotproduct

func solution(a []int, b []int) int {

	var total int
	for i, v := range a {
		total += v * b[i]
	}
	return total
}
