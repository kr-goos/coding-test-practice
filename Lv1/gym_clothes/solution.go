package gymclothes

import (
	"sort"
)

const UNAVAILABLE = -1

func solution(n int, lost []int, reserve []int) int {
	sort.Ints(lost)
	sort.Ints(reserve)

	answer := n - len(lost)
	for i, l := range lost {
		for j, r := range reserve {
			if l == r {
				answer++
				lost[i] = UNAVAILABLE
				reserve[j] = UNAVAILABLE
				break
			}
		}
	}

	for _, l := range lost {
		if l == UNAVAILABLE {
			continue
		}

		for j, r := range reserve {
			if r == UNAVAILABLE {
				continue
			}

			if l-1 == r || l+1 == r {
				answer++
				reserve[j] = UNAVAILABLE
				break
			}
		}

	}
	return answer
}
