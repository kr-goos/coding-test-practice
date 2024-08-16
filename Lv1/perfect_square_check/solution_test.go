package perfectsquarecheck

import (
	"testing"
)

const TESTCOUNT = 2

var (
	n      = [TESTCOUNT]int64{121, 3}
	result = [TESTCOUNT]int64{144, -1}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		v := solution(n[i])
		if v != result[i] {
			t.Errorf("[testcase %d failure] result : %d solution result : %d", i, result[i], v)
		}
	}
}
