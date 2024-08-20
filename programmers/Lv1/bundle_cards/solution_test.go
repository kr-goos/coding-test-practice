package bundlecards

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	const COUNT = 2
	card1TestCases := [COUNT][]string{
		{"i", "drink", "water"},
		{"i", "water", "drink"},
	}
	card2TestCases := [COUNT][]string{
		{"want", "to"},
		{"want", "to"},
	}
	goalTestCases := [COUNT][]string{
		{"i", "want", "to", "drink", "water"},
		{"i", "want", "to", "drink", "water"},
	}

	successResults := [COUNT]string{
		"Yes",
		"No",
	}

	for i := 0; i < COUNT; i++ {
		v := solution(card1TestCases[i], card2TestCases[i], goalTestCases[i])
		if v == successResults[i] {
			fmt.Printf("testcase %d success\n", i+1)
		} else {
			t.Fatalf("testcase %d failure\n", i+1)
		}
	}

}
