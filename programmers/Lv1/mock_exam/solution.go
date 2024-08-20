package mockexam

func solution(answers []int) []int {
	pattern1 := []int{1, 2, 3, 4, 5}
	pattern2 := []int{2, 1, 2, 3, 2, 4, 2, 5}
	pattern3 := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

	correct := make([]int, 3)
	p1len, p2len, p3len := len(pattern1), len(pattern2), len(pattern3)
	for i, answer := range answers {
		if pattern1[i%p1len] == answer {
			correct[0]++
		}

		if pattern2[i%p2len] == answer {
			correct[1]++
		}

		if pattern3[i%p3len] == answer {
			correct[2]++
		}
	}

	var max int
	for _, c := range correct {
		if c > max {
			max = c
		}
	}

	result := make([]int, 0, 3)
	for i, c := range correct {
		if c == max {
			result = append(result, i+1)
		}
	}

	return result
}
