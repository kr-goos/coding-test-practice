package sortnumbersdescending

import (
	"testing"
)

var (
	n      int64 = 118372
	result int64 = 873211
)

func TestSolution(t *testing.T) {
	answer := solution(n)
	if answer != result {
		t.Errorf("answer %d != result %d", answer, result)
	}
}
