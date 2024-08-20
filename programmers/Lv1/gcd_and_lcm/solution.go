package gcdandlcm

func solution(n int, m int) []int {
	gcd := Gcd(n, m)

	return []int{gcd, n * m / gcd}
}

func Gcd(n, m int) int {
	if n < m {
		n, m = m, n
	}

	if m == 0 {
		return n
	}
	return Gcd(m, n%m)
}
