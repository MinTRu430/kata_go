package utils

import "strconv"

var romanNumber = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}

type arabicNumber struct {
	Value  int
	Symbol string
}

func RomanToInt(s string) int {
	total := 0
	prevValue := 0

	_, err := strconv.Atoi(s)
	if err == nil {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	for i := len(s) - 1; i >= 0; i-- {
		value := romanNumber[s[i]]
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}
	return total
}

func IntToRoman(num int) string {
	if num <= 0 {
		panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
	}

	var romanNumerals = []arabicNumber{
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

	result := ""
	for _, numeral := range romanNumerals {
		for num >= numeral.Value {
			result += numeral.Symbol
			num -= numeral.Value
		}
	}

	return result
}
