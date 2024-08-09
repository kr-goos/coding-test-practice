package calculateshortfall

func solution(price int, money int, count int) int64 {

	var total int
	for i := 1; i <= count; i++ {
		total += i * price
	}

	if shortfall := int64(total - money); shortfall > 0 {
		return shortfall
	}

	return 0
}

func improvedSolution(price int, money int, count int) int64 {

	total := price * count * (count + 1) / 2

	if total > money {
		return int64(total - money)

	}

	return 0
}
