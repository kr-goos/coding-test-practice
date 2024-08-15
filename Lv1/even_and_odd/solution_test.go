package evenandodd

import "testing"

func TestSolution(t *testing.T) {
	result := []string{"Odd", "Even"}

	for i, v := range []int{3, 4} {
		if result[i] != solution(v) {
			t.Errorf("test %d failure", i)
		}
	}
}
