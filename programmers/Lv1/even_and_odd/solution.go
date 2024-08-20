package evenandodd

const (
	EVEN = "Even"
	ODD  = "Odd"
)

func solution(num int) string {
	if num%2 == 0 {
		return EVEN
	}
	return ODD
}
