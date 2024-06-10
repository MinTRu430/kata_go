package utils

import (
	"regexp"
	"strconv"
)

func IsDigit(s string) (int, bool) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	if num < 1 || num > 10 {
		panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более.")
	}
	return num, true
}

var romanNumeralPattern = regexp.MustCompile("^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$")

func IsRomanDigit(s string) (int, bool) {
	_, err := strconv.Atoi(s)
	if err == nil {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	if !romanNumeralPattern.MatchString(s) {
		panic("паника неправильный ввод римских символов")
	}

	num := RomanToInt(s)
	if num < 1 || num > 10 {
		panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более.")
	}
	return num, true
}
