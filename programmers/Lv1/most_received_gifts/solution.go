package mostreceivedgifts

import "strings"

func solution(friends []string, gifts []string) int {

	m := initializeGiftExchangeMap(friends)

	setupGiftExchangeMap(m, gifts)
	giftsToReceive := make(map[string]int)

	for i := 0; i < len(friends); i++ {
		for j := i + 1; j < len(friends); j++ {

			if m[friends[i]][friends[j]] > m[friends[j]][friends[i]] {
				giftsToReceive[friends[i]]++
			} else if m[friends[i]][friends[j]] < m[friends[j]][friends[i]] {
				giftsToReceive[friends[j]]++
			} else {
				if m[friends[i]]["index"] > m[friends[j]]["index"] {
					giftsToReceive[friends[i]]++
				} else if m[friends[i]]["index"] < m[friends[j]]["index"] {
					giftsToReceive[friends[j]]++
				} else {
					continue
				}
			}
		}
	}

	var max int
	for _, v := range giftsToReceive {
		if v > max {
			max = v
		}
	}

	return max
}

func initializeGiftExchangeMap(f []string) map[string]map[string]int {
	m := map[string]map[string]int{}

	for _, v1 := range f {
		m[v1] = make(map[string]int)
		for _, v2 := range f {
			m[v1][v2] = 0
		}
	}

	return m
}

func setupGiftExchangeMap(m map[string]map[string]int, g []string) {
	for _, gift := range g {
		s := strings.Split(gift, " ")
		// s[0] : 준 선물
		// s[1] : 받은 선물-
		m[s[0]][s[1]]++
		m[s[0]]["index"]++
		m[s[1]]["index"]--
	}
}
