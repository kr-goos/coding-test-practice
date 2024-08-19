package customsort

func solution(strings []string, n int) []string {

	for i := 0; i < len(strings); i++ {
		for j := i + 1; j < len(strings); j++ {
			if strings[i][n] > strings[j][n] {
				strings[i], strings[j] = strings[j], strings[i]
			} else if strings[i][n] == strings[j][n] {
				if strings[i] > strings[j] {
					strings[i], strings[j] = strings[j], strings[i]
				}
			}
		}
	}

	return []string{}
}
