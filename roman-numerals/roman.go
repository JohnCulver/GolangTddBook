package romannumerals

import "strings"

type RomanNumeral struct {
	Value  uint16
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
func ConvertToRomanOld(num uint16) string {

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

func ConvertToRoman(num uint16) string {

	var result strings.Builder

	for _, rn := range allRomanNumerals {
		for num >= rn.Value {
			result.WriteString(rn.Symbol)
			num -= rn.Value
		}
	}

	return result.String()
}

// This was my first attempt at writing the for loop
// it is less efficient
//
// Deprecated: Use ConvertToArabic
func ConvertToArabicOld(symbolInput string) uint16 {
	var result uint16
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

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
