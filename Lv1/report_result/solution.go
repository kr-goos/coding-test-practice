package reportresult

import (
	"strings"
)

func solution(id_list []string, report []string, k int) []int {
	reportedCountMap := make(map[string]int)
	reportedIdMap := make(map[string]map[string]struct{})
	for _, ids := range report {
		s := strings.Split(ids, " ")

		m, ok := reportedIdMap[s[0]]
		if !ok {
			m = make(map[string]struct{})
		} else {
			_, exist := m[s[1]]
			if exist {
				continue
			}
		}
		m[s[1]] = struct{}{}
		reportedIdMap[s[0]] = m
		reportedCountMap[s[1]]++
	}

	suspendedId := map[string]struct{}{}

	for id, count := range reportedCountMap {
		if count >= k {
			suspendedId[id] = struct{}{}
		}
	}

	answer := make([]int, len(id_list))
	for i, id := range id_list {
		for reportedId, _ := range reportedIdMap[id] {
			if _, ok := suspendedId[reportedId]; ok {
				answer[i]++
			}
		}
	}

	return answer
}
