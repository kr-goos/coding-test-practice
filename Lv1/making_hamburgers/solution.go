package makinghamburgers

import (
	"strconv"
	"strings"
)

const (
	BREAD      = "1"
	VEGETABLE  = "2"
	PATTY      = "3"
	HAMBBERGER = BREAD + VEGETABLE + PATTY + BREAD
)

func solution(ingredient []int) int {

	var ingredientStr strings.Builder
	for _, v := range ingredient {
		ingredientStr.WriteString(strconv.Itoa(v))
	}

	var count int
	v := ingredientStr.String()
	for {
		i := strings.Index(v, HAMBBERGER)
		if i == -1 {
			break
		}
		count++
		v = v[:i] + v[i+len(HAMBBERGER):]
	}

	return count
}
