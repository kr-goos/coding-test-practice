package sortstringdescending

import "testing"

var (
	s      = "Zbcdefg"
	result = "gfedcbZ"
)

func TestSolution(t *testing.T) {
	answer := solution(s)
	if answer != result {
		t.Errorf("result %s != solution result %s", answer, result)
	}
}

func TestSimpleSolution(t *testing.T) {
	answer := simpleSolution(s)
	if answer != result {
		t.Errorf("result %s != solution result %s", answer, result)
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(s)
	}
}

func BenchmarkSimpleSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleSolution(s)
	}
}
