package summissingnumbers

func solution(numbers []int) int {

	m := make(map[int]struct{}, 9)
	for _, num := range numbers {
		m[num] = struct{}{}
	}

	var answer int
	for i := 1; i <= 9; i++ {
		if _, ok := m[i]; !ok {
			answer += i
		}
	}

	return answer
}
