package main

import (
	"strings"
)

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := arabic; i > 0; i-- {
		result.WriteString("I")
	}

	return result.String()
}
