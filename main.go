package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Калькулятор запущен. Введите выражение (например, 3 + 2, I + V, римские числа должны быть написаны заглавными буквами):")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Калькулятор завершает работу.")
			return
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		if strings.ContainsAny(input, "IVXLCDM") {
			romanResult := arabicToRoman(result)
			fmt.Println("Результат:", romanResult)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func evaluateExpression(expression string) (int, error) {
	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		panic("неверное количество аргументов")
	}

	firstOperand, err := parseOperand(tokens[0])
	if err != nil {
		return 0, err
	}

	operator := tokens[1]

	secondOperand, err := parseOperand(tokens[2])
	if err != nil {
		return 0, err
	}

	// Проверяем, что операнды имеют одинаковый тип (арабские или римские)
	if (isRomanNumeral(tokens[0]) && !isRomanNumeral(tokens[2])) || (!isRomanNumeral(tokens[0]) && isRomanNumeral(tokens[2])) {
		panic("операнды должны быть одного типа (арабские или римские)")
	}

	if (firstOperand < 1 || firstOperand > 10) && (secondOperand < 1 || secondOperand > 10) {
		panic("неверные операнды")
	}

	var result int
	switch operator {
	case "+":
		result = firstOperand + secondOperand
	case "-":
		result = firstOperand - secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "/":
		result = firstOperand / secondOperand
	default:
		panic("неверная операция")
	}

	return result, nil
}

func parseOperand(operandStr string) (int, error) {
	arabicNums := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	if arabicNum, ok := arabicNums[operandStr]; ok {
		return arabicNum, nil
	}

	num, err := strconv.Atoi(operandStr)
	if err != nil {
		panic("неверный формат числа")
	}

	if num < 1 || num > 10 {
		panic("число должно быть от 1 до 10 включительно")
	}

	return num, nil
}

func isRomanNumeral(str string) bool {
	for _, r := range str {
		if !strings.ContainsRune("IVXLCDM", r) {
			return false
		}
	}
	return true
}

func arabicToRoman(num int) string {
	if num < 1 {
		panic("число должно быть больше нуля")
	}
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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

	var result strings.Builder
	for _, numeral := range romanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}
