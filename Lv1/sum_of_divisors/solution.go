package sumofdivisors

func solution(n int) int {
	var answer int
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			answer += i
		}
	}
	return answer + n
}
