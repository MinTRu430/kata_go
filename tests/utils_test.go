package utils

import (
	"kata_go/utils"
	"testing"
)

type IsDigitTest struct {
	input     string
	expected  int
	wantPanic bool
}

func TestIsDigit(t *testing.T) {
	tests := []IsDigitTest{
		{"5", 5, true},
		{"10", 10, false},
		{"0", 0, false},
		{"11", 0, false},
		{"a", 0, false},
	}

	for _, tt := range tests {
		if tt.wantPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%s, not panic", tt.input)
				}
			}()
		}
		result, valid := utils.IsDigit(tt.input)
		if !tt.wantPanic && (!valid || result != tt.expected) {
			t.Errorf("IsDigit %s = %d, %t; want %d, true", tt.input, result, valid, tt.expected)
		}
	}
}

type IsRomanDigitTest struct {
	input     string
	expected  int
	wantPanic bool
}

func TestIsRomanDigit(t *testing.T) {
	tests := []IsRomanDigitTest{
		{"I", 1, false},
		{"IV", 4, false},
		{"X", 10, false},
		{"XL", 0, true},
		{"a", 0, true},
	}

	for _, tt := range tests {
		if tt.wantPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf(" %s, not panic", tt.input)
				}
			}()
		}
		result, valid := utils.IsRomanDigit(tt.input)
		if !tt.wantPanic && (!valid || result != tt.expected) {
			t.Errorf("IsDigit %s = %d, %t; want %d, true", tt.input, result, valid, tt.expected)
		}
	}
}
