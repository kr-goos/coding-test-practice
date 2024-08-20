package reportresult

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	id_list = [TESTCOUNT][]string{
		{"muzi", "frodo", "apeach", "neo"},
		{"con", "ryan"},
	}
	report = [TESTCOUNT][]string{
		{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"},
		{"ryan con", "ryan con", "ryan con", "ryan con"},
	}
	k      = [TESTCOUNT]int{2, 3}
	result = [TESTCOUNT][]int{
		{2, 1, 1, 0},
		{0, 0},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("id_list : %v\n", id_list[i])
		fmt.Printf("report : %v\n", report[i])
		fmt.Printf("k : %d\n", k[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(id_list[i], report[i], k[i]))
	}
}
