package runningrace

// 파라미터를 그대로 이용한 방법
func solution(players []string, callings []string) []string {
	for _, calledName := range callings {
	LOOP:
		for idx, player := range players {
			if player != calledName {
				continue
			}
			players[idx-1], players[idx] = player, players[idx-1]
			break LOOP
		}
	}
	return players
}

// 파라미터의 자료구조를 map 으로 변형한 방법
func solution2(players []string, callings []string) []string {

	playerMap := stringSliceToIndexMap(players)

	for _, calledPlayer := range callings {
		currRank := playerMap[calledPlayer]
		players[currRank], players[currRank-1] = players[currRank-1], players[currRank]
		playerMap[calledPlayer]--
		playerMap[players[currRank]]++
	}
	return players
}

func stringSliceToIndexMap(arr []string) map[string]int {
	m := make(map[string]int, len(arr))
	for i, v := range arr {
		m[v] = i
	}
	return m
}
