package coke

func solution(a int, b int, n int) int {
	var count int
	for n >= a {
		returnedCoke := n / a * b
		remainedCoke := n % a

		n = remainedCoke + returnedCoke
		count += returnedCoke
	}

	return count
}
