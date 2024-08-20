package personalitytest

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	survey = [TESTCOUNT][]string{
		{"AN", "CF", "MJ", "RT", "NA"},
		{"TR", "RT", "TR"},
	}
	choices = [TESTCOUNT][]int{
		{5, 3, 2, 7, 5},
		{7, 1, 3},
	}
	result = [TESTCOUNT]string{
		"TCMA",
		"RCJA",
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("survey : %v\n", survey[i])
		fmt.Printf("choices : %v\n", choices[i])
		fmt.Printf("result : %s\n", result[i])
		fmt.Printf("solution result : %s\n", solution(survey[i], choices[i]))
	}
}
