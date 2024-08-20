package bundlecards

func solution(cards1 []string, cards2 []string, goal []string) string {

	for _, v := range goal {
		if len(cards1) > 0 && cards1[0] == v {
			cards1 = cards1[1:]
			continue
		}

		if len(cards2) > 0 && cards2[0] == v {
			cards2 = cards2[1:]
			continue
		}
		return "No"
	}

	return "Yes"
}
