package romannumerals

import (
	"testing"
	"testing/quick"
)

var cases = []struct {
	Num uint16
	Rn  string
}{
	{Num: 1, Rn: "I"},
	{Num: 2, Rn: "II"},
	{Num: 3, Rn: "III"},
	{Num: 4, Rn: "IV"},
	{Num: 5, Rn: "V"},
	{Num: 6, Rn: "VI"},
	{Num: 7, Rn: "VII"},
	{Num: 8, Rn: "VIII"},
	{Num: 9, Rn: "IX"},
	{Num: 10, Rn: "X"},
	{Num: 14, Rn: "XIV"},
	{Num: 18, Rn: "XVIII"},
	{Num: 20, Rn: "XX"},
	{Num: 39, Rn: "XXXIX"},
	{Num: 40, Rn: "XL"},
	{Num: 47, Rn: "XLVII"},
	{Num: 49, Rn: "XLIX"},
	{Num: 50, Rn: "L"},
	{Num: 100, Rn: "C"},
	{Num: 90, Rn: "XC"},
	{Num: 400, Rn: "CD"},
	{Num: 500, Rn: "D"},
	{Num: 900, Rn: "CM"},
	{Num: 1000, Rn: "M"},
	{Num: 1984, Rn: "MCMLXXXIV"},
	{Num: 3999, Rn: "MMMCMXCIX"},
	{Num: 2014, Rn: "MMXIV"},
	{Num: 1006, Rn: "MVI"},
	{Num: 798, Rn: "DCCXCVIII"},
}

func TestConversionToRoman(t *testing.T) {

	for _, testCase := range cases {
		got := ConvertToRoman(testCase.Num)
		want := testCase.Rn
		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	}

}

func TestConversionToArabic(t *testing.T) {
	for _, testCase := range cases {
		got := ConvertToArabic(testCase.Rn)
		want := testCase.Num
		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	}
}

func TestPropertiesOfConverstion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)

		return fromRoman == arabic
	}

	err := quick.Check(assertion, nil)
	if err != nil {
		t.Error("failed checks", err)
	}
}
