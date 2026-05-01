package iteration

func Repeat(char string, repetitions uint) string {
	var repeated string

	for range repetitions {
		repeated += char
	}

	return repeated
}
