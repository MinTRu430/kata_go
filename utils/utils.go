package utils

import "strconv"

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

func IsRomanDigit(s string) (int, bool) {
	num := RomanToInt(s)
	if num < 1 || num > 10 {
		panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более.")
	}
	return num, true
}
