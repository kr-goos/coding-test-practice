package roughkeyboard

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	keymapTestCases := [][]string{
		{"ABACD", "BCEFD"},
		{"AA"},
		{"AGZ", "BSSS"},
	}
	targetsTestCases := [][]string{
		{"ABCD", "AABB"},
		{"B"},
		{"ASA", "BGZ"},
	}
	successResults := [][]int{
		{9, 4},
		{-1},
		{4, 6},
	}

	for i := 0; i < 3; i++ {
		v := solution(keymapTestCases[i], targetsTestCases[i])

		fmt.Printf("keymapTestCases[%d] : %v\n", i, keymapTestCases[i])
		fmt.Printf("targetsTestCases[%d] : %v\n", i, targetsTestCases[i])
		for j, sr := range successResults[i] {
			if v[j] == sr {
				fmt.Printf("success result : %d ,solution result : %d\n", sr, v[j])
			} else {
				t.Fatalf("success result : %d ,solution result : %d\n", sr, v[j])
			}
		}

	}

}
