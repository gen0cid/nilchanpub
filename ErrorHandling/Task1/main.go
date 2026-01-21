package main

import (
	"errors"
	"fmt"
)

func subtraction(a, b float64) (float64, error) {

	return a - b, nil
}

func addition(a, b float64) (float64, error) {

	return a + b, nil
}

func multiplication(a, b float64) (float64, error) {

	return a * b, nil
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на 0")
	}
	return a / b, nil
}
func validateArgs(a, b float64) error {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return errors.New("аргументы вне допустимого диапазона [-1000; 1000]")
	}
	return nil
}

func performOperation(op string, a, b float64) {
	err := validateArgs(a, b)
	if err != nil {
		fmt.Printf("Ошибка: %s | Операция: %s | Аргументы: %.2f, %.2f\n", err.Error(), op, a, b)
		return
	}
	var res float64
	var error error
	switch op {
	case "+":
		res, error = addition(a, b)
	case "-":
		res, error = subtraction(a, b)
	case "*":
		res, error = multiplication(a, b)
	case "/":
		res, error = division(a, b)
	default:
		fmt.Printf("Неизвестная операция: %s\n", op)
		return
	}

	if error != nil {
		fmt.Printf("Ошибка: %s | Операция: %s | Аргументы: %.2f, %.2f\n", error.Error(), op, a, b)
	} else {
		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", a, op, b, res)
	}
}

func main() {

}
