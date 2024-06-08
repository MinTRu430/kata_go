package utils

import (
	"kata_go/utils"
	"testing"
)

type RomanToIntTest struct {
	input    string
	expected int
}

func TestRomanToInt(t *testing.T) {
	tests := []RomanToIntTest{
		{"I", 1},
		{"IV", 4},
		{"IX", 9},
		{"X", 10},
		{"XL", 40},
		{"XC", 90},
		{"C", 100},
	}

	for _, tt := range tests {
		result := utils.RomanToInt(tt.input)
		if result != tt.expected {
			t.Errorf("RomanToInt %s = %d; want %d", tt.input, result, tt.expected)
		}
	}
}

type IntToRomanTest struct {
	input     int
	expected  string
	wantPanic bool
}

func TestIntToRoman(t *testing.T) {
	tests := []IntToRomanTest{
		{1, "I", false},
		{4, "IV", false},
		{9, "IX", false},
		{10, "X", false},
		{40, "XL", false},
		{90, "XC", false},
		{100, "C", false},
		{0, "", true},
		{-1, "", true},
	}

	for _, tt := range tests {
		if tt.wantPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf(" %d,not panic", tt.input)
				}
			}()
		}
		result := utils.IntToRoman(tt.input)
		if result != tt.expected {
			t.Errorf("IntToRoman %d = %s; want %s", tt.input, result, tt.expected)
		}
	}
}
