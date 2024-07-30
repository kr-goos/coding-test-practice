package smallpartialstring

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	const TESTCOUNT = 3
	tTestcases := [TESTCOUNT]string{
		"3141592",
		"500220839878",
		"10203",
	}

	pTestcases := [TESTCOUNT]string{
		"271",
		"7",
		"15",
	}

	successResults := [TESTCOUNT]int{
		2,
		8,
		3,
	}

	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("t : ", tTestcases[i])
		fmt.Println("p : ", pTestcases[i])
		fmt.Println("success result : ", successResults[i])
		fmt.Println("solution() result : ", solution(tTestcases[i], pTestcases[i]))
	}

}
