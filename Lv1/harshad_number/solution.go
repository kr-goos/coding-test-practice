package harshadnumber

import "strconv"

func solution(x int) bool {
	s := strconv.Itoa(x)
	var total int
	for i := 0; i < len(s); i++ {
		v, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		total += v
	}

	return x%total == 0
}

func solution2(x int) bool {
	temp := x
	var total int
	for temp > 0 {
		total += temp % 10
		temp /= 10
	}

	return x%total == 0
}
