package middlechar

func solution(s string) string {
	length := len(s)
	if length == 1 {
		return s
	}

	middle := length / 2
	if length%2 == 0 {
		return string(s[middle-1 : middle+1])
	}

	return string(s[middle])
}
