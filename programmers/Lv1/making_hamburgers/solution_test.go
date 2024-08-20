package makinghamburgers

import (
	"fmt"
	"testing"
)

const (
	TESTCOUNT = 2
)

var (
	ingredient = [][]int{
		{2, 1, 1, 2, 3, 1, 2, 3, 1},
		{1, 3, 2, 1, 2, 1, 3, 1, 2},
	}

	result = []int{
		2, 0,
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("ingredient[i] : ", ingredient[i])
		fmt.Println("result	 : ", result[i])
		fmt.Println("my solution result : ", solution(ingredient[i]))
	}
}
