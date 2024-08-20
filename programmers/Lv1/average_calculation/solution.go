package averagecalculation

func solution(arr []int) float64 {
	var total int
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}
