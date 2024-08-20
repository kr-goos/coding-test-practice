package year2016

import (
	"fmt"
	"time"
)

func solution(a int, b int) string {
	m := map[int]string{
		0: "SUN",
		1: "MON",
		2: "TUE",
		3: "WED",
		4: "THU",
		5: "FRI",
		6: "SAT",
	}
	date, err := time.Parse("2006-01-02", fmt.Sprintf("2016-%02d-%02d", a, b))
	if err != nil {
		return ""
	}

	return m[int(date.Weekday())]
}
