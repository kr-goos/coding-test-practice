package countandsumdivisors

func solution(left int, right int) int {
	var answer int
	for i := left; i <= right; i++ {
		v := countDivisors(i)
		if v == 0 {
			continue
		}

		if v%2 == 0 {
			answer += i
		} else {
			answer -= i
		}
	}
	return answer
}

func countDivisors(v int) int {
	var count int
	for i := 1; i <= v; i++ {
		if v%i == 0 {
			count++
		}
	}

	return count
}
