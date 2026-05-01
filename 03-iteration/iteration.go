package iteration

import "strings"

func Repeat(char string, repetitions uint) string {
	var result strings.Builder

	for range repetitions {
		result.WriteString(char)
	}

	return result.String()
}
