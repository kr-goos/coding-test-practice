package privacypolicyretentionperiod

import (
	"strconv"
	"strings"
	"time"
)

func solution(today string, terms []string, privacies []string) []int {
	termsMap := make(map[string]int, len(terms))
	for _, term := range terms {
		s := strings.Split(term, " ")
		retentionPeriod, err := strconv.Atoi(s[1])
		if err != nil {
			return nil
		}
		termsMap[s[0]] = retentionPeriod
	}

	answer := []int{}
	for i, privacy := range privacies {

		s := strings.Split(privacy, " ")

		if dateToDays(s[0])+termsMap[s[1]]*28-1 < dateToDays(today) {
			answer = append(answer, i+1)
		}
	}

	return answer
}

func dateToDays(date string) int {
	t, err := time.Parse("2006.01.02", date)
	if err != nil {
		return -1
	}

	yDays := t.Year() * 12 * 28
	mDays := int(t.Month()) * 28
	return yDays + mDays + t.Day()
}
