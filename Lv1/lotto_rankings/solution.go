package lottorankings

func solution(lottos []int, win_nums []int) []int {

	m := make(map[int]struct{})
	for _, num := range win_nums {
		m[num] = struct{}{}
	}

	var matchingCount, zeroCount int
	for _, num := range lottos {
		if num == 0 {
			zeroCount++
			continue
		}

		if _, ok := m[num]; ok {
			matchingCount++
		}
	}

	return []int{getRanking(matchingCount + zeroCount), getRanking(matchingCount)}
}

func getRanking(v int) int {
	switch v {
	case 6:
		return 1
	case 5:
		return 2
	case 4:
		return 3
	case 3:
		return 4
	case 2:
		return 5
	default:
		return 6
	}
}
