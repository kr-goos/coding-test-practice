package halloffame1

import "sort"

func solution(k int, score []int) []int {

	hallOfFame := make([]int, 0, k+1)
	result := make([]int, len(score))
	for i, s := range score {
		hallOfFame = append(hallOfFame, s)
		sort.Ints(hallOfFame)

		if len(hallOfFame) > k {
			hallOfFame = hallOfFame[1:]
		}
		result[i] = hallOfFame[0]
	}

	return result
}
