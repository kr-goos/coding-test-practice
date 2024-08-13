package personalitytest

import (
	"math"
	"strings"
)

const (
	// negative
	RYAN   = "R"
	CON    = "C"
	JAYG   = "J"
	APEACH = "A"

	// positive
	TUBE  = "T"
	FRODO = "F"
	MUZI  = "M"
	NEO   = "N"
)

func solution(survey []string, choices []int) string {

	m := make(map[string]int)

	for i, types := range survey {
		if choices[i] == 4 {
			continue
		}

		m[string(types[choices[i]/4])] += int(math.Abs(float64(4 - choices[i])))
	}

	negatives := [4]string{RYAN, CON, JAYG, APEACH}
	positives := [4]string{TUBE, FRODO, MUZI, NEO}

	var sb strings.Builder
	for i := 0; i < 4; i++ {
		if m[negatives[i]] >= m[positives[i]] {
			sb.WriteString(negatives[i])
		} else {
			sb.WriteString(positives[i])
		}
	}

	return sb.String()
}
