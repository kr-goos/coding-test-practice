package averagecalculation

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	arr = [TESTCOUNT][]int{
		{1, 2, 3, 4},
		{5, 5},
	}
	result = [TESTCOUNT]float64{2.5, 5}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("arr : %v\n", arr[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(arr[i]))
	}
}
