package overpainting

func solution(n int, m int, section []int) int {

	count, position := 0, section[0]

	for _, v := range section {
		if position > v {
			continue
		}
		position = v + m
		count++
	}

	return count
}
