package numberpair

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(X string, Y string) string {

	var xArr *[10]int = CharacterFrequencyArr(X)
	var yArr *[10]int = CharacterFrequencyArr(Y)

	var sb strings.Builder
	for i := 9; i >= 0; i-- {

		if xArr[i] <= 0 || yArr[i] <= 0 {
			continue
		}

		var min int = 3000001

		if xArr[i] < min {
			min = xArr[i]
		}

		if yArr[i] < min {
			min = yArr[i]
		}

		if i == 0 && sb.String() == "" && min != 0 {
			return "0"
		}

		for j := 0; j < min; j++ {
			fmt.Fprintf(&sb, "%d", i)
		}

	}

	if sb.String() == "" {
		return "-1"
	}

	return sb.String()
}

func CharacterFrequencyArr(s string) *[10]int {
	var arr [10]int

	for _, c := range s {
		v, _ := strconv.Atoi(string(c))
		arr[v]++
	}

	return &arr
}
