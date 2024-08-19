package year2016

import "testing"

func TestSolution(t *testing.T) {
	a := 5
	b := 24
	result := "TUE"
	r := solution(a, b)
	if r != result {
		t.Errorf("result : %s != %s", result, r)
	}
}
