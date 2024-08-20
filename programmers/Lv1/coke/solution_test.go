package coke

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	a      = [TESTCOUNT]int{2, 3}
	b      = [TESTCOUNT]int{1, 1}
	n      = [TESTCOUNT]int{20, 20}
	result = [TESTCOUNT]int{19, 9}
)

func TestSolution(t *testing.T) {

	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("a[i] : ", a[i])
		fmt.Println("b[i] : ", b[i])
		fmt.Println("n[i] : ", n[i])
		fmt.Println("result[i] : ", result[i])
		fmt.Println("my solution result : ", solution(a[i], b[i], n[i]))
	}
}
