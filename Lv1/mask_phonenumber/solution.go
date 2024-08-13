package maskphonenumber

import "strings"

func solution(phone_number string) string {

	var sb strings.Builder
	for i := 0; i < len(phone_number)-4; i++ {
		sb.WriteString("*")
	}

	sb.WriteString(phone_number[len(phone_number)-4:])

	return sb.String()
}
