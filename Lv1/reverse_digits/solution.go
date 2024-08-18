package reversedigits

func solution(n int64) []int {
	var answer []int
	for n > 0 {
		answer = append(answer, int(n%10))
		n /= 10
	}
	return answer
}
