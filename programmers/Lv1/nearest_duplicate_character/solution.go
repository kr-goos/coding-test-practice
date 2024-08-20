package nearestduplicatecharacter

func solution(s string) []int {

	result := make([]int, len(s))
	lastIndexMap := make(map[rune]int, len(s))
	for i, char := range s {
		if lastIdx, found := lastIndexMap[char]; found {
			result[i] = i - lastIdx
		} else {
			result[i] = -1
		}
		lastIndexMap[char] = i
	}

	return result
}
