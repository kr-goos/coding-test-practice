package numberswithinterval

func solution(x int, n int) []int64 {

	answer := make([]int64, 0, n)
	for i := 1; i <= n; i++ {
		answer = append(answer, int64(x*i))
	}
	return answer
}
