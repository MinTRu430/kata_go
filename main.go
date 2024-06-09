package main

import (
	"bufio"
	"fmt"
	"kata_go/utils"
	"os"
	"strings"
)

func stringToInt(s string) {
	x := strings.Split(s, " ")
	if len(x) < 3 {
		panic("Выдача паники, так как строка не является математической операцией.")
	} else if len(x) > 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	num1 := strings.TrimSpace(x[0])
	operator := strings.TrimSpace(x[1])
	num2 := strings.TrimSpace(x[2])

	if n1, ok := utils.IsDigit(num1); ok {
		if n2, ok := utils.IsDigit(num2); ok {
			fmt.Println(utils.Calculation(n1, n2, operator))
		} else {
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		}
	} else if r1, ok := utils.IsRomanDigit(num1); ok {
		if r2, ok := utils.IsRomanDigit(num2); ok {
			result := utils.Calculation(r1, r2, operator)
			if result <= 0 {
				panic("Выдача паники, так как в римской системе нет отрицательных чисел. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна выдать панику.")
			}
			fmt.Println(utils.IntToRoman(result))
		} else {
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		}
	} else {
		panic("Выдача паники, числа не являются корректными цифрами или римскими цифрами.")
	}

}

func normalizeInput(input string) string {
	input = strings.ReplaceAll(input, "+", " + ")
	input = strings.ReplaceAll(input, "-", " - ")
	input = strings.ReplaceAll(input, "*", " * ")
	input = strings.ReplaceAll(input, "/", " / ")
	input = strings.TrimSpace(input)
	return strings.Join(strings.Fields(input), " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	i := 0
	for {
		i++
		fmt.Printf("Пример %d:\n", i)
		input, _ := reader.ReadString('\n')

		stringToInt(normalizeInput(input))
	}
}
