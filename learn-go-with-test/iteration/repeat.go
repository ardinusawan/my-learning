package iteration

import "strings"

// Repeat get character and return repeated character for 5 times.
func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
