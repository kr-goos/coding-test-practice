package weirdstring

import "testing"

func TestSolution(t *testing.T) {
	s := "try hello world"
	result := "TrY HeLlO WoRlD"
	r := solution(s)
	if r != result {
		t.Errorf("%s != %s", result, r)
	}
}
