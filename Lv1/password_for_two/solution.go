package passwordfortwo

const (
	Z = 'z'
)

func solution(s string, skip string, index int) string {

	m := make(map[rune]struct{})
	for _, char := range skip {
		m[char] = struct{}{}
	}

	var answer string
	for _, char := range s {
		for i := 0; i < index; i++ {
			char += 1
			if char > Z {
				char = 'a'
			}

			if _, ok := m[char]; ok {
				i--
			}
		}
		answer += string(char)
	}

	return answer
}
