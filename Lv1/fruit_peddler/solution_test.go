package fruitpeddler

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	const TESTCOUNT = 2
	k := [TESTCOUNT]int{3, 4}
	m := [TESTCOUNT]int{4, 3}
	score := [TESTCOUNT][]int{
		{1, 2, 3, 1, 2, 3, 1},
		{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2},
	}
	success := [TESTCOUNT]int{8, 33}

	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("k[i] : ", k[i])
		fmt.Println("m[i] : ", m[i])
		fmt.Println("score[i] : ", score[i])
		fmt.Println("success[i] : ", success[i])
		fmt.Println("solution result : ", solution(k[i], m[i], score[i]))
	}
}
