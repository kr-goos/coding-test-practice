package main

// 파라미터를 그대로 이용한 방법
func solution(name []string, yearning []int, photo [][]string) []int {

	memScoreArr := make([]int, len(photo))

	for i := range photo {
		for _, personName := range photo[i] {
			idx := findIndexInStringArray(name, personName)
			if idx == -1 {
				continue
			}
			memScoreArr[i] += yearning[idx]
		}
	}

	return memScoreArr
}

func findIndexInStringArray(strArr []string, target string) int {

	for i, str := range strArr {
		if str == target {
			return i
		}
	}

	return -1
}

// 파라미터의 자료구조를 map 으로 변형한 방법
func solution2(name []string, yearning []int, photo [][]string) []int {

	infoMap := make(map[string]int, len(name))
	for i, n := range name {
		infoMap[n] = yearning[i]
	}

	memScoreArr := make([]int, len(photo))
	for i := range photo {
		for _, personName := range photo[i] {
			memScoreArr[i] += infoMap[personName]
		}
	}

	return memScoreArr
}
