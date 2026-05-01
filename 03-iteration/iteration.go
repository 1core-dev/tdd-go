package iteration

const repeatCount = 6

func Repeat(char string) string {
	var repeated string

	for range repeatCount {
		repeated += char
	}

	return repeated
}
