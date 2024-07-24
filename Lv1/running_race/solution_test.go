package runningrace

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	playersTestCases := []string{
		"mumu", "soe", "poe", "kai", "mine",
	}
	callingsTestCases := []string{
		"kai", "kai", "mine", "mine",
	}

	successResult := []string{
		"mumu", "kai", "mine", "soe", "poe",
	}

	results := solution(playersTestCases, callingsTestCases)
	for i := range results {
		if results[i] != successResult[i] {
			t.Fatalf("result : %s != sucecssResult : %s\n", results[i], successResult[i])
		}
		fmt.Printf("result : %s == sucecssResult : %s\n", results[i], successResult[i])
	}
}

func TestSolution2(t *testing.T) {
	playersTestCases := []string{
		"mumu", "soe", "poe", "kai", "mine",
	}
	callingsTestCases := []string{
		"kai", "kai", "mine", "mine",
	}

	successResult := []string{
		"mumu", "kai", "mine", "soe", "poe",
	}

	results := solution2(playersTestCases, callingsTestCases)
	for i := range results {
		if results[i] != successResult[i] {
			t.Fatalf("result : %s != sucecssResult : %s\n", results[i], successResult[i])
		}
		fmt.Printf("result : %s == sucecssResult : %s\n", results[i], successResult[i])
	}

}

func BenchmarkSolution(b *testing.B) {

	for i := 0; i < b.N; i++ {
		playersTestCases := []string{
			"mumu", "soe", "poe", "kai", "mine",
		}
		callingsTestCases := []string{
			"kai", "kai", "mine", "mine",
		}
		_ = solution(playersTestCases, callingsTestCases)
	}
}

func BenchmarkSolution2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		playersTestCases := []string{
			"mumu", "soe", "poe", "kai", "mine",
		}
		callingsTestCases := []string{
			"kai", "kai", "mine", "mine",
		}
		_ = solution2(playersTestCases, callingsTestCases)
	}
}
