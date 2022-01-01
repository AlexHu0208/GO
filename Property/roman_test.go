package property_based

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic int
	Want   string
}{
	{Arabic: 1, Want: "I"},
	{Arabic: 2, Want: "II"},
	{Arabic: 3, Want: "III"},
	{Arabic: 4, Want: "IV"},
	{Arabic: 5, Want: "V"},
	{Arabic: 6, Want: "VI"},
	{Arabic: 7, Want: "VII"},
	{Arabic: 8, Want: "VIII"},
	{Arabic: 9, Want: "IX"},
	{Arabic: 10, Want: "X"},
	{Arabic: 14, Want: "XIV"},
	{Arabic: 18, Want: "XVIII"},
	{Arabic: 20, Want: "XX"},
	{Arabic: 39, Want: "XXXIX"},
	{Arabic: 40, Want: "XL"},
	{Arabic: 47, Want: "XLVII"},
	{Arabic: 49, Want: "XLIX"},
	{Arabic: 50, Want: "L"},
	{Arabic: 100, Want: "C"},
	{Arabic: 90, Want: "XC"},
	{Arabic: 400, Want: "CD"},
	{Arabic: 500, Want: "D"},
	{Arabic: 900, Want: "CM"},
	{Arabic: 1000, Want: "M"},
	{Arabic: 1984, Want: "MCMLXXXIV"},
	{Arabic: 3999, Want: "MMMCMXCIX"},
	{Arabic: 2014, Want: "MMXIV"},
	{Arabic: 1006, Want: "MVI"},
	{Arabic: 798, Want: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Want), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Want, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Want)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfCoversion(t *testing.T) {
	assertion := func(arabic int) bool {
		if arabic < 0 || arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
