package collatzconjecture

func solution(num int) int {
	var count int
	if num == 1 {
		return 0
	}

	for {
		if count >= 500 {
			return -1
		}

		count++

		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}

		if num == 1 {
			return count
		}
	}
}
