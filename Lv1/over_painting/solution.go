package overpainting

func solution(n int, m int, section []int) int {

	count, position := 0, section[0]

	for i := range section {
		if position > section[i] {
			continue
		}
		position = section[i] + m
		count++
	}

	return count
}
