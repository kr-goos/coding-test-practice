package foodfightcompetition

import (
	"strconv"
	"strings"
)

func solution(food []int) string {

	var leftStr string
	temp := []int{}
	for i, c := range food {
		if i == 0 {
			continue
		}

		for j := 0; j < c/2; j++ {
			temp = append(temp, i)
			leftStr += strconv.Itoa(i)
		}
	}

	var rightStr string
	for j := len(temp) - 1; j >= 0; j-- {
		rightStr += strconv.Itoa(temp[j])
	}

	return leftStr + "0" + rightStr
}

func ImprovedSolution(food []int) string {
	var leftBuilder strings.Builder

	for i, c := range food {
		if i == 0 {
			continue
		}

		str := strconv.Itoa(i)
		for j := 0; j < c/2; j++ {
			leftBuilder.WriteString(str)
		}
	}

	return leftBuilder.String() + "0" + reverseString(leftBuilder.String())
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
