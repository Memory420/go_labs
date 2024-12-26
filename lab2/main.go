package main

import (
	"errors"
	"fmt"
)

func helloFunc(name string) string {
	return fmt.Sprintf("Привет, %s!", name)
}

func printEvenNums(a, b int64) error {
	if a > b {
		return errors.New("левая граница диапазона больше правой")
	}

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}

func main() {
	fmt.Println(helloFunc("Иван"))

	err := printEvenNums(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	err = printEvenNums(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	result, err := apply(3, 5, "+")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f\n", result)
	}

	result, err = apply(7, 10, "*")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f\n", result)
	}

	result, err = apply(3, 5, "#")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f\n", result)
	}

	result, err = apply(10, 0, "/")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f\n", result)
	}
}
