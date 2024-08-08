package threemusketeers

func solution(number []int) int {
	var count int
	for i := 0; i < len(number); i++ {
		for j := i + 1; j < len(number); j++ {
			for k := j + 1; k < len(number); k++ {
				if number[i]+number[j]+number[k] == 0 {
					count++
				}
			}
		}
	}
	return count
}
