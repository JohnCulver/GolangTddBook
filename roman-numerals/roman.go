package romannumerals

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// This was my first attempt at writing the for loop
// it is less efficient
//
// Deprecated: Use ConvertToRoman
func ConvertToRomanOld(num int) string {

	var result strings.Builder

	for num > 0 {
		for _, rn := range allRomanNumerals {
			if num >= rn.Value {
				result.WriteString(rn.Symbol)
				num -= rn.Value
				break
			}
		}
	}

	return result.String()
}

func ConvertToRoman(num int) string {

	var result strings.Builder

	for _, rn := range allRomanNumerals {
		for num >= rn.Value {
			result.WriteString(rn.Symbol)
			num -= rn.Value
		}
	}

	return result.String()
}

func ConvertToArabic(symbolInput string) int {
	var result int
	for len(symbolInput) > 0 {
		for _, rn := range allRomanNumerals {
			if strings.HasPrefix(symbolInput, rn.Symbol) {
				result += rn.Value
				symbolInput = strings.TrimPrefix(symbolInput, rn.Symbol)
			}
		}
	}

	return result
}
