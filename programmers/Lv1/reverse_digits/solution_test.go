package reversedigits

import "testing"

func TestSolution(t *testing.T) {
	var n int64 = 12345
	result := []int{5, 4, 3, 2, 1}

	sr := solution(n)
	for i := range result {
		if result[i] != sr[i] {
			t.Errorf("result : %d != solution result : %d", result[i], sr[i])
		}
	}
}
