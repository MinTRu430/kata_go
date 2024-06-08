package utils

import (
	"kata_go/utils"
	"testing"
)

type CalculationTest struct {
	x         int
	y         int
	operator  string
	expected  int
	wantPanic bool
}

func TestCalculation(t *testing.T) {
	tests := []CalculationTest{
		{5, 3, "+", 8, false},
		{5, 3, "-", 2, false},
		{5, 3, "*", 15, false},
		{6, 3, "/", 2, false},
		{6, 0, "/", 0, true},
		{6, 3, "%", 0, true},
	}

	for _, tt := range tests {
		if tt.wantPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%d %s %d, not panic", tt.x, tt.operator, tt.y)
				}
			}()
		}
		result := utils.Calculation(tt.x, tt.y, tt.operator)
		if result != tt.expected {
			t.Errorf("Calculation %d, %d, %s  = %d; want %d", tt.x, tt.y, tt.operator, result, tt.expected)
		}
	}
}
