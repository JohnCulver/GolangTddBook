package romannumerals

import "testing"

func TestConversion(t *testing.T) {
	testSet := []struct {
		num int
		rn  string
	}{
		{1, "I"},
		{2, "II"},
		{4, "IV"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{24, "XXIV"},
	}

	for _, testCase := range testSet {
		got := ConvertToRoman(testCase.num)
		want := testCase.rn
		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	}

}
