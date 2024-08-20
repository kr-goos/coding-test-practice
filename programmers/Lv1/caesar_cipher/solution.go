package caesarcipher

func solution(s string, n int) string {

	runes := []rune(s)
	offset := rune(n)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = (r-'a'+offset)%26 + 'a'
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = (r-'A'+offset)%26 + 'A'
		}
	}

	return string(runes)
}
