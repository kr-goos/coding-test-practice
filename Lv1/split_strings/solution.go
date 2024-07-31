package slpitstrings

func solution(s string) int {

	var result int
	for i := 0; i < len(s); {
		var count, othercount int
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				count++
			} else {
				othercount++
			}

			if count == othercount {
				i = j + 1
				result++
				break
			}
		}

		if count != othercount {
			result++
			break
		}
	}

	return result
}
