package babbling2

import "strings"

const (
	AYA = "aya"
	YE  = "ye"
	WOO = "woo"
	MA  = "ma"
)

func solution(babbling []string) int {

	var count int
	allowed := []string{AYA, YE, WOO, MA}
	for _, word := range babbling {

		if isRepeatPattern(word, allowed) {
			continue
		}

		for _, pattern := range allowed {
			word = strings.ReplaceAll(word, pattern, " ")
			if strings.TrimSpace(word) == "" {
				count++
				break
			}
		}

	}

	return count
}

func isRepeatPattern(s string, patterns []string) bool {

	for _, pattern := range patterns {
		if strings.Contains(s, pattern+pattern) {
			return true
		}
	}

	return false
}

// 60점 짜리 실패 코드
// 실패 사유 : 공백으로 변경으로 인한 yayae 와같은 문자열에서
// y'aya'e  를 제거하면 ye 가 되기 때문이다.
// ye 도 발음이 가능한 문자이다.
// 하지만, 위 예문에서는 발음이 가능한 단어들의 연속이어야 한다고 했다.
func solution2(babbling []string) int {

	var count int
	allowed := []string{AYA, YE, WOO, MA}
	for _, word := range babbling {

		if isRepeatPattern(word, allowed) {
			continue
		}

		for _, pattern := range allowed {
			word = strings.ReplaceAll(word, pattern, "")
			if word == "" {
				count++
				break
			}
		}
	}

	return count
}
