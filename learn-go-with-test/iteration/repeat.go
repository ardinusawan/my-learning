package iteration

const repeatCount = 5

// Repeat get character and return repeated character for 5 times.
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
