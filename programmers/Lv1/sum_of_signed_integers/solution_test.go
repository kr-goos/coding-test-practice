package sumofsignedintegers

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	absolutes = [TESTCOUNT][]int{
		{4, 7, 12},
		{1, 2, 3},
	}
	signs = [TESTCOUNT][]bool{
		{true, false, true},
		{false, false, true},
	}
	result = [TESTCOUNT]int{9, 0}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("absolutes : %v\n", absolutes[i])
		fmt.Printf("signs : %v\n", signs[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(absolutes[i], signs[i]))

	}
}
