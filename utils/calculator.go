package utils

func Calculation(x, y int, operator string) int {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		if y == 0 {
			panic("Выдача паники, на ноль делить нельзя")
		}
		return x / y
	default:
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
}
