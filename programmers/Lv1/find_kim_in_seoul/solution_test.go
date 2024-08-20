package findkiminseoul

import "testing"

func TestSolution(t *testing.T) {
	seoul := []string{
		"Jane", "Kim",
	}
	result := "김서방은 1에 있다"

	r := solution(seoul)
	if r != result {
		t.Errorf("result %s != solution result %s", result, r)
	}
}
