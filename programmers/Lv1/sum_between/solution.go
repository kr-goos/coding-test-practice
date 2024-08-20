package sumbetween

func solution(a int, b int) int64 {
	if a > b {
		a, b = b, a
	}
	var answer int
	for i := a; i <= b; i++ {
		answer += i
	}
	return int64(answer)
}
