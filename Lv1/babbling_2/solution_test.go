package babbling2

import "testing"

const (
	TESTCOUNT = 3
)

var (
	babbling = [TESTCOUNT][]string{
		{"aya", "yee", "u", "maa"},
		{"ayaye", "uuu", "yeye", "yemawoo", "ayaayaa"},
		{"mayaa"},
	}
	expected = [TESTCOUNT]int{1, 2, 0}
)

func TestSolution(t *testing.T) {

	for i := 0; i < TESTCOUNT; i++ {
		result := solution(babbling[i])
		if result != expected[i] {
			t.Errorf("For input %v, expected %d but got %d", babbling[i], expected[i], result)
		}
	}
}

func TestSolution2(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		result := solution2(babbling[i])
		if result != expected[i] {
			t.Errorf("For input %v, expected %d but got %d", babbling[i], expected[i], result)
		}
	}
}
