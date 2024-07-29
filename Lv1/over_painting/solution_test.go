package overpainting

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	nTestCases := []int{8, 5, 4}
	mTestCases := []int{4, 4, 1}
	sectionTestCases := [][]int{
		{2, 3, 6},
		{1, 3},
		{1, 2, 3, 4},
	}
	successResult := []int{2, 1, 4}

	for i := 0; i < 3; i++ {
		v := solution(nTestCases[i], mTestCases[i], sectionTestCases[i])
		if v == successResult[i] {
			fmt.Printf("%d test success\n", i+1)
			continue
		}
		t.Fatalf("%d test failure\n", i+1)
	}
}
