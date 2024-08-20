package smallpartialstring

import "strconv"

func solution(t string, p string) int {
	plen := len(p)
	tlen := len(t)

	pValue, err := strconv.Atoi(p)
	if err != nil {
		return -1
	}

	var count int
	for i := 0; i < tlen-plen+1; i++ {
		tPartialValue, err := strconv.Atoi(t[i : i+plen])
		if err != nil {
			return -1
		}

		if tPartialValue <= pValue {
			count++
		}
	}
	return count
}
