package romannumerals

import "strings"

func ConvertToRoman(num int) string {
	var result strings.Builder

	for num > 0 {
		switch {
		case num >= 10:
			result.WriteString("X")
			num -= 10
		case num >= 9:
			result.WriteString("IX")
			num -= 9
		case num >= 4:
			result.WriteString("IV")
			num -= 4

		default:
			result.WriteString("I")
			num -= 1
		}
	}

	return result.String()
}
