package roughkeyboard

func solution(keymap []string, targets []string) []int {
	m := make(map[string]int)
	for _, keyArr := range keymap {
		for i := 0; i < len(keyArr); i++ {
			key := string(keyArr[i])
			v, ok := m[key]
			if !ok || v > i {
				m[key] = i + 1
				continue
			}
		}
	}

	answer := make([]int, len(targets))
	for i, targetStr := range targets {
		for j := 0; j < len(targetStr); j++ {
			v, ok := m[string(targetStr[j])]
			if !ok {
				answer[i] = -1
				break
			}
			answer[i] += v
		}
	}

	return answer
}
